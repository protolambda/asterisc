// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// RISCVMetaData contains all meta data concerning the RISCV contract.
var RISCVMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_preimageOracle\",\"type\":\"address\",\"internalType\":\"contractIPreimageOracle\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"preimageOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPreimageOracle\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"step\",\"inputs\":[{\"name\":\"stateData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"localContext\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051620021ff380380620021ff83398101604081905261003191610056565b600080546001600160a01b0319166001600160a01b0392909216919091179055610086565b60006020828403121561006857600080fd5b81516001600160a01b038116811461007f57600080fd5b9392505050565b61216980620000966000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80633acc02841461003b578063e14ced3214610085575b600080fd5b60005461005b9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100986100933660046120bf565b6100a6565b60405190815260200161007c565b6000610370565b8060005260206000fd5b905090565b600067ffffffffffffffff6100cf565b90565b909116919050565b60006100f2601f6100ed63ffffffff5b85610287565b610110565b92915050565b6000600160405b1b905090565b60006001603f6100ff565b60006001831b821680156101345767ffffffffffffffff841c841b8317915061014d565b61014a67ffffffffffffffff85603f031c841690565b91505b5092915050565b600061015e610105565b82168015610190577fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000083179150610194565b8291505b50919050565b60006101ae6101a76100f8565b8484010690565b9392505050565b60006101ae6101c26100f8565b8484030690565b60006101ae8383026100bc565b60006101ae8383046100bc565b60006101ae6101f184610154565b6101fa84610154565b056100bc565b60008282066101ae565b60006101ae61021884610154565b61022184610154565b076100bc565b60006100f282196100bc565b60008282106101ae565b60008282116101ae565b60006101ae61025584610154565b61025e84610154565b1290565b60006101ae61027084610154565b61027984610154565b1390565b60008183146101ae565b60008183166101ae565b60008183176101ae565b60008183186101ae565b60006101ae83831b6100bc565b600082821c6101ae565b60006101ae6102ca84610154565b831d6100bc565b60208110156101945760ff83168260081b1791508260081c92506001810190506102d1565b6000602060005b01905090565b600060206102fd6102f6565b600060086102fd610303565b600060086102fd61030f565b600060016102fd61031b565b600060016102fd610327565b600060086102fd610333565b600060086102fd61033f565b600060086102fd61034b565b60006101006102fd610357565b60806040511461037f57600080fd5b6084861461038c57600080fd5b610394610363565b6020870335146103a357600080fd5b60206103bd6103b0610363565b601f808216602003160190565b87010184146103cb57600080fd5b61022484146103d957600080fd5b6103e1610363565b6080016040526103ef610363565b86608037611455565b6080015160209190910360031b1c90565b60800180517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60209390930360031b92831b19169290911b919091179052565b50565b60006100b760206103f86102f6565b6104498160206104096102f6565b60006100b760086103f8610303565b610449816008610409610303565b60006100b760086103f861030f565b61044981600861040961030f565b60006100b760016103f8610327565b6104c0600180610409610327565b565b60006100b760016103f861031b565b61044981600161040961031b565b60006100b760086103f8610333565b610449816008610409610333565b60006100b760086103f861033f565b61044981600861040961033f565b60006100b760086103f861034b565b61044981600861040961034b565b6000610543601f8361023d565b156105545761055462bad4e96100ad565b6105726105626008846101c9565b61056d6100cc610357565b61019a565b608081015160c01c6101ae565b80610588575050565b610594601f5b8261023d565b156105a5576105a562bad4e96100ad565b6105b36105626008836101c9565b60808101805177ffffffffffffffffffffffffffffffffffffffffffffffff1660c085901b179052505050565b505050565b60006105ef6104a3565b600181146105ff57600391505090565b6106076104c2565b801561061e57600181146106275760029250505090565b60009250505090565b60019250505090565b610638610363565b6080a06000610645610363565b60802090506106526105e5565b60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff919091161790565b60006100f2600b6100ed8460146102b2565b6102b2565b60006100f2600b6100ed6106b3601f6106ae8760076102b2565b610287565b6106cd6106c18760196102b2565b60056102a5565b6102a5565b610291565b60006100f2600c6100ed61070e6106f46106ed87601f6102b2565b600c6102a5565b6106cd61070760016106ae8a60076102b2565b600b6102a5565b6106cd6107246106c1603f6106ae8a60196102b2565b6106cd610737600f6106ae8b60086102b2565b60016102a5565b60006100f260136100ed84600c6102b2565b60006100f260136100ed61078561077261076b87601f6102b2565b60136102a5565b6106cd61070760ff6106ae8a600c6102b2565b6106cd6107a261079b60016106ae8a60146102b2565b600a6102a5565b6106cd6103ff6106ae8a60156102b2565b60006100f2607f5b83610287565b60006100f2601f6106ae8460076102b2565b60006100f260076106ae84600c6102b2565b60006100f2601f6106ae84600f6102b2565b60006100f2601f6106ae8460146102b2565b60006100f28260196102b2565b60006100f28260146102b2565b600061083a6020610835603c856101c9565b6101c9565b90506100f26102245b8261019a565b6000610855601f6107bb565b156108675761086763bad10ad06100ad565b61087083610823565b803561087e60205b8361019a565b915061088b8460056102b2565b8160005b603b8110156108f95784356108a560208761019a565b95506108b660016106ae86856102b2565b80156108c957600181146108de576108ef565b600084815260208390526040902093506108ef565b600082815260208590526040902093505b505060010161088f565b5060805193508381146109135761091363badf00d16100ad565b509095945050505050565b61092a601f5b82610287565b1561093c5761093c63bad10ad06100ad565b61094583610823565b826109506020610878565b915061095d8360056102b2565b60005b603b8110156109cb57833561097760205b8661019a565b945061098860016106ae85856102b2565b801561099b57600181146109b0576109c1565b600085815260208390526040902094506109c1565b600082815260208690526040902094505b5050600101610960565b50506109d681608052565b5050505050565b600060088311156109f5576109f563bad512e06100ad565b610a076107bb601f610227565b610227565b610a146100cc8783610849565b610a1e82856101b5565b6000610a41610a2d601f610227565b6106ae610a3b60018b6101b5565b8961019a565b610a5388610a4e85602082565b6101b5565b6000610a66610a62848961027d565b1590565b15610aa35760ff8d03610a8057610a8063bad222206100ad565b610a8d6100cc8e85610849565b935060009150610aa08a610a4e87604082565b90505b85610ab26100cc8460036102a5565b1c955083610ac46100cc8360036102a5565b1c9350505050600091505b86821015610b5957610af7610ae4601f610227565b6106ae84610a4e6001610a4e8d8d61019a565b6000610b03868361027d565b60018114610b16578015610b2657610b32565b600886901c9560ff169150610b32565b600884901c9360ff1691505b50610b4a610b3f826100bc565b6106cd8960086102a5565b96505050600182019150610acf565b505050508315610b8157610b736001610a4e8560036102a5565b610b7d8183610110565b9150505b95945050505050565b60008060008084610b9b878761019a565b60005b6040811015610c635780610bb3602082610233565b8060018114610bc7578015610bd857610be5565b8760081b97508960081b9950610be5565b8660081b96508860081b98505b50610c07610bf38584610233565b6106ae6000610c028987610233565b61027d565b15610c595760ff8c610c266100cc610c1f8e876101b5565b60036102a5565b1c168160018114610c3c578015610c4b57610c56565b9781179760ff8b179a50610c56565b9681179660ff8a1799505b50505b5050600101610b9e565b50505093509350935093565b6020821115610c8557610c8563bad512e16100ad565b610c92610924601f610227565b610cb3610c9f601f610227565b6106ae610cad6001876101b5565b8561019a565b610cbd82846101b5565b610cc8868683610b8a565b9350610cd76100cc8b88610849565b8319168117610ce78b828961091e565b5050610cf3848661027d565b15610d045750505050505050505050565b60ff8a03610d1957610d1963bad222216100ad565b610d266100cc8b86610849565b9019169091179250610d41905087610d3b8490565b8361091e565b50505050505050565b6109d68585858585610c6f565b60008360018114610d7f5760028114610d845760038114610d9557610d7f63badc0de06100ad565b610da4565b610d8e8483610291565b9350610da4565b610da16107bb85610227565b93505b509392505050565b6000610db8601f6107bb565b610dc38160206101b5565b610dcd818661023d565b15610dd6578094505b50610dee6100cc6001610de984876101b5565b610849565b80610dfd6100cc8460036102a5565b1b905080610e126100cc610c1f8860206101b5565b1c90508460031b9150610e2361044c565b821b8117610e308161045b565b505050610e43610e3e600090565b610478565b5090919050565b600080805460405163e03110e160e01b8152846004820152856024820152620186a06040600060448460008786f1925050508015610e92576000519250602051915050610ea0565b50610ea063badf00d06100ad565b9250929050565b6000610eb161044c565b610eb9610469565b8160f81c60018103610f2157604080516000858152336020528983526060902091527effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01000000000000000000000000000000000000000000000000000000000000001792505b50610f2c8183610e4a565b925082610f3f57600093505050506101ae565b610f4b601f5b86610287565b610f568160206101b5565b610f60818961023d565b15610f69578097505b50610f74848861023d565b15610f7d578396505b610f8b610c1f8860206101b5565b6001901b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01199350610fc26100cc8260036102a5565b93841c9382901c9150610fd5878461019a565b9250610fe083610478565b610fef6001610de983896101b5565b84191684831617925061100d600184611008848a6101b5565b61091e565b50949695505050505050565b6110236011610536565b80605d811461111c57605e811461111c5760d6811461113b5760de811461116057603f81146111d9576040811461126f57601981146112c9576038811461137757607b811461139557607c811461139557607181146113a9576087811461139557608481146113955760b2811461139557608681146113955760dc81146113db5760a381146113e55760e981146113955760148114611395576015811461139557603b811461139557604e811461139557604f81146113955760a081146113955760d78114611395576101168114611395576101058114611448576101a681146114485760658114611448576105e063f001ca116100ad565b611126600a610536565b61113260ff82166104d1565b506105e06104b2565b6111476001601e6102a5565b61115281600a61057f565b506105e060005b600b61057f565b61116a600a610536565b611174600b610536565b81806111cb57611185610fff6107bb565b80156111a3576111a061119a826110006101b5565b8461019a565b92505b506111ac6104fc565b6111b781600a61057f565b6111c96111c4848361019a565b61050b565b505b5050506105e0611159600090565b6111e3600a610536565b6111ed600b610536565b6111f7600c610536565b6000808480156112275760038114611232576005811461123c5767ffffffffffffffff5b9250604d5b915061124e565b600092506000611220565b8392506000611220565b611247898587610ea7565b9250600091505b5061125a82600a61057f565b61126581600b61057f565b5050505050505050565b611279600a610536565b611283600b610536565b61128d600c610536565b60008084600181146112325760028114611232576004811461123257600681146112bf5767ffffffffffffffff61121b565b6112478486610dac565b6112d3600a610536565b6112dd600b610536565b60008082600381146112fd5767ffffffffffffffff925060169150611360565b8480156113485760018114611352576002811461135257600381146113485760048114611352576005811461134857600681146113595767ffffffffffffffff9350604d925061135e565b60005b935061135e565b600161134b565b600193505b505b5061136c82600a61057f565b610d4181600b61057f565b61138b67ffffffffffffffff5b600a61057f565b6105e0600d611159565b61139f6000611384565b6105e06000611159565b6113b3600b610536565b682a00000000000005396113ce600260018360105b86610c6f565b505061139f611384600090565b61139f6001611384565b6113ef600a610536565b6113f9600b610536565b81600781146114135761140e620f00126100ad565b611440565b61142c60026001690400000000000000040060106113c8565b6114366000611384565b6114406000611159565b505050505050565b6105e063f001ca116100ad565b61145d6104a3565b156114735761146a610630565b60005260206000f35b611488611483600161056d6104df565b6104ee565b611490610486565b6114a060ff6000806004856109dd565b6114a9816107b3565b6114b2826107c1565b6114bb836107d3565b6114c4846107e5565b6114cd856107f7565b6114d686610809565b856003811461157157602381146115e7576063811461163d576013811461172257601b811461182a57603381146118e757603b8114611aee5760378114611ce15760178114611d0157606f8114611d1d5760678114611d5d5760738114611db157602f8114611e5157600f811461205257600781146120525760278114612052576053811461205f5761156c63f001c0de6100ad565b61206c565b61157a8861067d565b611589610a6260045b88610287565b61159a60016106c860035b8a610287565b6115a387610536565b6115b2610843600b5b86610110565b9350506115c4600260018484876109dd565b925050506115d2818861057f565b5061156c6115e260045b8b61019a565b610495565b6115f088610694565b6115fb6001876102a5565b61160485610536565b61160d87610536565b61161a610843600b6115ac565b93505061162c60026001838587610d4a565b50505061156c6115e26115dc600490565b61164684610536565b61164f84610536565b6000878015611685576001811461169657600481146116a857600581146116b257600681146116c457600781146116ce576116e3565b61168f838561027d565b91506116e3565b61168f60016106ae610a02868861027d565b61168f8385610247565b61168f60016106ae610a028688610247565b61168f8385610233565b6116e060016106ae610a028688610233565b91505b50808015611707576116f48c6106d2565b9350611700848e61019a565b9c50611715565b61171260048e61019a565b9c505b5050505061156c89610495565b61172b84610536565b6117348961067d565b600087801561177a576001811461178b576002811461179957600381146117a357600481146117ad57600581146117b7576006811461180857600781146118125761181f565b611784838561019a565b915061181f565b611784846106c8603f610f45565b6117848385610247565b6117848385610233565b611784838561029b565b6117c28360066102b2565b80156117d557601081146117ec57611802565b6117e58561068f603f5b87610287565b9250611802565b6117ff856117fa603f6117df565b6102bc565b92505b5061181f565b6117848385610291565b61181c8385610287565b91505b5061162c818a61057f565b61183384610536565b61183c8961067d565b600087801561185a576001811461186c576005811461187d5761181f565b611784611867848661019a565b6100d7565b611784611867856106c8601f6117df565b611888601f84610287565b6118938460056102b2565b80156118a657602081146118c9576118da565b6118c2601f5b6100ed6118bc63ffffffff611594565b856102b2565b93506118da565b6118d76118ac83601f6101b5565b93505b50505061162c818a61057f565b6118f084610536565b6118f984610536565b600084600181146119f0578880156119485760018114611980576002811461198e576003811461199857600481146119a257600581146119ac57600681146119dc57600781146119e657611802565b86801561195c576020811461196d5761197a565b611966858761019a565b935061197a565b61197785876101b5565b93505b50611802565b6117e5856106c8603f6117df565b6117e58486610247565b6117e58486610233565b6117e5848661029b565b8680156119c057602081146119ce5761197a565b6119668661068f603f611583565b611977866117fa603f611583565b6117e58486610291565b6117ff8486610287565b888015611a345760018114611a455760028114611a635760038114611a6c5760048114611a765760058114611aa15760068114611ab25760078114611acc57611ae2565b611a3e84866101c9565b9250611ae2565b611a3e611a5185610154565b611a5a87610154565b0260401c6100bc565b611a3e84611a51565b611a3e8486611a5a565b838015611a8e57611a8785876101e3565b9350611a9b565b67ffffffffffffffff5b93505b50611ae2565b838015611a8e57611a8785876101d6565b838015611ac357611a87858761020a565b85935050611ae2565b838015611add576118c28587610200565b859350505b505061162c818a61057f565b611af784610536565b611b0084610536565b60008460018114611be157888015611b275760018114611b7b5760058114611b8c57611802565b868015611b3b5760208114611b5b5761197a565b611966611867611b4e63ffffffff611583565b61056d63ffffffff611594565b611977611867611b6e63ffffffff611583565b610a4e63ffffffff611594565b6117e5611867866106c8601f611583565b611b96601f6100e7565b878015611baa5760208114611bc857611bd9565b611bc1601f5b6100ed6118bc63ffffffff8b610287565b9450611bd9565b611bd6611bb083601f6101b5565b94505b50505061181f565b888015611c0d5760048114611c2f5760058114611c535760068114611c7f5760078114611cac57611ae2565b611a3e611867611c2063ffffffff6117df565b61083563ffffffff5b89610287565b838015611a8e57611a87611867611c45876100d7565b611c4e896100d7565b6101e3565b838015611a8e57611a87611867611c6d63ffffffff611583565b611c7a63ffffffff611594565b6101d6565b838015611ca357611a87611867611c95876100d7565b611c9e896100d7565b61020a565b611a98866100d7565b838015611cd8576118c2611867611cc663ffffffff611583565b611cd363ffffffff611594565b610200565b6118d7866100d7565b611cea8861073e565b611cf581600c6102a5565b90506115d2818861057f565b611d0a8861073e565b611cf56115dc601f6100ed84600c6102a5565b611d2688610750565b611d3060046115dc565b611d3a818961057f565b50611d576115e2611d5160146100ed8560016102a5565b8c61019a565b5061206c565b611d6684610536565b611d6f8961067d565b611d796004611d51565b611d83818a61057f565b50611daa6115e2611d9d600167ffffffffffffffff61029b565b6106ae610971600b6115ac565b505061206c565b848015611e1657611dc189610816565b85611dcf610a626004611594565b15611de057611ddd87610536565b90505b611dea6003611c29565b611df5818385610d57565b92505050611e03818961057f565b50611e116115e26004611d51565b611d57565b611e218960146102b2565b8015611e3b57611e366115e260045b8d61019a565b611daa565b611e448d611019565b611daa6115e26004611e30565b611e5c6001866102a5565b611e66600861058e565b611e71600483610233565b1715611e8357611e8362bada706100ad565b611e8c85610536565b611e978460026102b2565b8060028114611fcd5760038114611ff557611eb187610536565b611ebc60048661027d565b15611ecd57611eca816100d7565b90505b611edc600260018088886109dd565b80848015611f365760018114611f475760048114611f4f5760088114611f5957600c8114611f635760108114611f6d5760148114611f845760188114611f8e57601c8114611f9857611f31630f001a706100ad565b611fab565b611f40848461019a565b9250611fab565b839250611fab565b611f40848461029b565b611f408484610291565b611f408484610287565b611f778385610247565b15611f3157839250611fab565b611f778385610262565b611f778385610233565b611fa2838561023d565b15611fab578392505b50611fbb60036001848a8a610d4a565b611fc5818e61057f565b505050612040565b611fdc600260018087876109dd565b611fe6818c61057f565b50611ff083610528565b612040565b6001612008612002610519565b8561027d565b1561202b5761201688610536565b61202560026001838989610d4a565b50600090505b612035818c61057f565b506120406000610528565b5050505061156c6115e26115dc600490565b61156c6115e260046115dc565b61206c6115e260046115dc565b50505050505050505061146a610630565b60008083601f84011261208f57600080fd5b50813567ffffffffffffffff8111156120a757600080fd5b602083019150836020828501011115610ea057600080fd5b6000806000806000606086880312156120d757600080fd5b853567ffffffffffffffff808211156120ef57600080fd5b6120fb89838a0161207d565b9097509550602088013591508082111561211457600080fd5b506121218882890161207d565b9699959850966040013594935050505056fea26469706673582212204bf07ebd5a7d3fef4e7f3c13d6576876bde01f9a5cf9bf3c1ad756d8155b2a4a64736f6c634300080f0033",
}

// RISCVABI is the input ABI used to generate the binding from.
// Deprecated: Use RISCVMetaData.ABI instead.
var RISCVABI = RISCVMetaData.ABI

// RISCVBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RISCVMetaData.Bin instead.
var RISCVBin = RISCVMetaData.Bin

// DeployRISCV deploys a new Ethereum contract, binding an instance of RISCV to it.
func DeployRISCV(auth *bind.TransactOpts, backend bind.ContractBackend, _preimageOracle common.Address) (common.Address, *types.Transaction, *RISCV, error) {
	parsed, err := RISCVMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RISCVBin), backend, _preimageOracle)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RISCV{RISCVCaller: RISCVCaller{contract: contract}, RISCVTransactor: RISCVTransactor{contract: contract}, RISCVFilterer: RISCVFilterer{contract: contract}}, nil
}

// RISCV is an auto generated Go binding around an Ethereum contract.
type RISCV struct {
	RISCVCaller     // Read-only binding to the contract
	RISCVTransactor // Write-only binding to the contract
	RISCVFilterer   // Log filterer for contract events
}

// RISCVCaller is an auto generated read-only Go binding around an Ethereum contract.
type RISCVCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RISCVTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RISCVTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RISCVFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RISCVFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RISCVSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RISCVSession struct {
	Contract     *RISCV            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RISCVCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RISCVCallerSession struct {
	Contract *RISCVCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RISCVTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RISCVTransactorSession struct {
	Contract     *RISCVTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RISCVRaw is an auto generated low-level Go binding around an Ethereum contract.
type RISCVRaw struct {
	Contract *RISCV // Generic contract binding to access the raw methods on
}

// RISCVCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RISCVCallerRaw struct {
	Contract *RISCVCaller // Generic read-only contract binding to access the raw methods on
}

// RISCVTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RISCVTransactorRaw struct {
	Contract *RISCVTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRISCV creates a new instance of RISCV, bound to a specific deployed contract.
func NewRISCV(address common.Address, backend bind.ContractBackend) (*RISCV, error) {
	contract, err := bindRISCV(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RISCV{RISCVCaller: RISCVCaller{contract: contract}, RISCVTransactor: RISCVTransactor{contract: contract}, RISCVFilterer: RISCVFilterer{contract: contract}}, nil
}

// NewRISCVCaller creates a new read-only instance of RISCV, bound to a specific deployed contract.
func NewRISCVCaller(address common.Address, caller bind.ContractCaller) (*RISCVCaller, error) {
	contract, err := bindRISCV(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RISCVCaller{contract: contract}, nil
}

// NewRISCVTransactor creates a new write-only instance of RISCV, bound to a specific deployed contract.
func NewRISCVTransactor(address common.Address, transactor bind.ContractTransactor) (*RISCVTransactor, error) {
	contract, err := bindRISCV(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RISCVTransactor{contract: contract}, nil
}

// NewRISCVFilterer creates a new log filterer instance of RISCV, bound to a specific deployed contract.
func NewRISCVFilterer(address common.Address, filterer bind.ContractFilterer) (*RISCVFilterer, error) {
	contract, err := bindRISCV(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RISCVFilterer{contract: contract}, nil
}

// bindRISCV binds a generic wrapper to an already deployed contract.
func bindRISCV(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RISCVMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RISCV *RISCVRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RISCV.Contract.RISCVCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RISCV *RISCVRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RISCV.Contract.RISCVTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RISCV *RISCVRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RISCV.Contract.RISCVTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RISCV *RISCVCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RISCV.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RISCV *RISCVTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RISCV.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RISCV *RISCVTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RISCV.Contract.contract.Transact(opts, method, params...)
}

// PreimageOracle is a free data retrieval call binding the contract method 0x3acc0284.
//
// Solidity: function preimageOracle() view returns(address)
func (_RISCV *RISCVCaller) PreimageOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RISCV.contract.Call(opts, &out, "preimageOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreimageOracle is a free data retrieval call binding the contract method 0x3acc0284.
//
// Solidity: function preimageOracle() view returns(address)
func (_RISCV *RISCVSession) PreimageOracle() (common.Address, error) {
	return _RISCV.Contract.PreimageOracle(&_RISCV.CallOpts)
}

// PreimageOracle is a free data retrieval call binding the contract method 0x3acc0284.
//
// Solidity: function preimageOracle() view returns(address)
func (_RISCV *RISCVCallerSession) PreimageOracle() (common.Address, error) {
	return _RISCV.Contract.PreimageOracle(&_RISCV.CallOpts)
}

// Step is a paid mutator transaction binding the contract method 0xe14ced32.
//
// Solidity: function step(bytes stateData, bytes proof, bytes32 localContext) returns(bytes32)
func (_RISCV *RISCVTransactor) Step(opts *bind.TransactOpts, stateData []byte, proof []byte, localContext [32]byte) (*types.Transaction, error) {
	return _RISCV.contract.Transact(opts, "step", stateData, proof, localContext)
}

// Step is a paid mutator transaction binding the contract method 0xe14ced32.
//
// Solidity: function step(bytes stateData, bytes proof, bytes32 localContext) returns(bytes32)
func (_RISCV *RISCVSession) Step(stateData []byte, proof []byte, localContext [32]byte) (*types.Transaction, error) {
	return _RISCV.Contract.Step(&_RISCV.TransactOpts, stateData, proof, localContext)
}

// Step is a paid mutator transaction binding the contract method 0xe14ced32.
//
// Solidity: function step(bytes stateData, bytes proof, bytes32 localContext) returns(bytes32)
func (_RISCV *RISCVTransactorSession) Step(stateData []byte, proof []byte, localContext [32]byte) (*types.Transaction, error) {
	return _RISCV.Contract.Step(&_RISCV.TransactOpts, stateData, proof, localContext)
}
