package challenger

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	op_challenger "github.com/ethereum-optimism/optimism/op-challenger"
	"github.com/ethereum-optimism/optimism/op-challenger/config"
	op_e2e_challenger "github.com/ethereum-optimism/optimism/op-e2e/e2eutils/challenger"
	"github.com/ethereum-optimism/optimism/op-service/metrics"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/stretchr/testify/require"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/log"
)

func WithAsterisc(
	t *testing.T,
	rollupCfg *rollup.Config,
	l2Genesis *core.Genesis,
	rollupEndpoint string,
	l2Endpoint string,
) op_e2e_challenger.Option {
	return func(c *config.Config) {
		c.TraceTypes = append(c.TraceTypes, config.TraceTypeAsterisc)
		c.RollupRpc = rollupEndpoint
		applyAsteriscConfig(c, t, rollupCfg, l2Genesis, l2Endpoint)
	}
}

func applyAsteriscConfig(
	c *config.Config,
	t *testing.T,
	rollupCfg *rollup.Config,
	l2Genesis *core.Genesis,
	l2Endpoint string,
) {
	require := require.New(t)
	c.L2Rpc = l2Endpoint
	root := op_e2e_challenger.FindMonorepoRoot(t)
	c.AsteriscBin = root + "rvgo/bin/asterisc"
	c.AsteriscServer = root + "rvsol/lib/optimism/op-program/bin/op-program"
	c.AsteriscAbsolutePreState = root + "rvgo/bin/prestate.json"
	c.AsteriscSnapshotFreq = 10_000_000

	genesisBytes, err := json.Marshal(l2Genesis)
	require.NoError(err, "marshall l2 genesis config")
	genesisFile := filepath.Join(c.Datadir, "l2-genesis.json")
	require.NoError(os.WriteFile(genesisFile, genesisBytes, 0o644))
	c.AsteriscL2GenesisPath = genesisFile

	rollupBytes, err := json.Marshal(rollupCfg)
	require.NoError(err, "marshall rollup config")
	rollupFile := filepath.Join(c.Datadir, "rollup.json")
	require.NoError(os.WriteFile(rollupFile, rollupBytes, 0o644))
	c.AsteriscRollupConfigPath = rollupFile
}

func NewChallenger(t *testing.T, ctx context.Context, sys op_e2e_challenger.EndpointProvider, name string, options ...op_e2e_challenger.Option) *op_e2e_challenger.Helper {
	log := testlog.Logger(t, log.LevelDebug).New("role", name)
	log.Info("Creating challenger")
	cfg := NewChallengerConfig(t, sys, options...)
	chl, err := op_challenger.Main(ctx, log, cfg)
	require.NoError(t, err, "must init challenger")
	require.NoError(t, chl.Start(ctx), "must start challenger")

	return op_e2e_challenger.NewHelper(log, t, require.New(t), cfg.Datadir, chl)
}

func NewChallengerConfig(t *testing.T, sys op_e2e_challenger.EndpointProvider, options ...op_e2e_challenger.Option) *config.Config {
	// Use the NewConfig method to ensure we pick up any defaults that are set.
	l1Endpoint := sys.NodeEndpoint("l1")
	l1Beacon := sys.L1BeaconEndpoint()
	cfg := config.NewConfig(common.Address{}, l1Endpoint, l1Beacon, t.TempDir())
	// The devnet can't set the absolute prestate output root because the contracts are deployed in L1 genesis
	// before the L2 genesis is known.
	cfg.AllowInvalidPrestate = true
	cfg.TxMgrConfig.NumConfirmations = 1
	cfg.TxMgrConfig.ReceiptQueryInterval = 1 * time.Second
	if cfg.MaxConcurrency > 4 {
		// Limit concurrency to something more reasonable when there are also multiple tests executing in parallel
		cfg.MaxConcurrency = 4
	}
	cfg.MetricsConfig = metrics.CLIConfig{
		Enabled:    true,
		ListenAddr: "127.0.0.1",
		ListenPort: 0, // Find any available port (avoids conflicts)
	}
	for _, option := range options {
		option(&cfg)
	}
	require.NotEmpty(t, cfg.TxMgrConfig.PrivateKey, "Missing private key for TxMgrConfig")
	require.NoError(t, cfg.Check(), "op-challenger config should be valid")

	if cfg.AsteriscBin != "" {
		_, err := os.Stat(cfg.AsteriscBin)
		require.NoError(t, err, "asterisc should be built. Make sure you've run make")
	}
	if cfg.AsteriscServer != "" {
		_, err := os.Stat(cfg.AsteriscServer)
		require.NoError(t, err, "op-program should be built. Make sure you've run make prestate")
	}
	if cfg.AsteriscAbsolutePreState != "" {
		_, err := os.Stat(cfg.AsteriscAbsolutePreState)
		require.NoError(t, err, "asterisc pre-state should be built. Make sure you've run make prestate")
	}
	if cfg.PollInterval == 0 {
		cfg.PollInterval = time.Second
	}

	return &cfg
}