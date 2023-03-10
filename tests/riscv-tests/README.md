# riscv-tests

This is a subset of test outputs generated by https://github.com/riscv-software-src/riscv-tests
The generated test vectors are checked into version control to make the testing more portable.

The test-vectors are licensed by the test generation 
[license](https://github.com/riscv-software-src/riscv-tests/blob/master/LICENSE),
also available here [`LICENSE`](./LICENSE)

Building unit tests (requires `riscv64-unknown-elf` toolchain):
```shell
git clone git@github.com:riscv-software-src/riscv-tests.git
cd riscv-tests

export RISCV=/opt/riscv
git submodule update --init --recursive
autoconf
./configure --prefix=$RISCV/target
RISCV_PREFIX="/opt/riscv/bin/riscv64-unknown-elf-" make

# files are output to the isa/ directory, e.g.
ls isa/rv64ui-p-*

# Selecting the test-vectors for Asterisc:
mkdir riscv-tests/rv64ua-p
mkdir riscv-tests/rv64ui-p
mkdir riscv-tests/rv64um-p
mkdir riscv-tests/benchmarks

cp isa/rv64ua-p-* riscv-tests/rv64ua-p/
cp isa/rv64ui-p-* riscv-tests/rv64ui-p/
cp isa/rv64um-p-* riscv-tests/rv64um-p/
cp benchmarks/*.riscv riscv-tests/benchmarks
cp benchmarks/*.riscv.dump riscv-tests/benchmarks
```

- `riscv_test.h` defines test environment things
- The "TVM" (test virtual machine) is the feature set required by a test
- We're only interested in `rv64u*`: **64** bit **u**ser-level integer-only instructions.
  - We only care about `i` (base integer set), and `a` (atomics), `m` (multiplication) extensions
  - We don't need the floating point, 32 bit, supervisor variants.
- And there are different target environments too. But we only care about single-core.
  - `p` = single core, physical memory
  - `v` = virtual memory enabled, may be interesting (TODO)
  - ignore multi-core tests

Test structure is documented here: https://riscv.org/wp-content/uploads/2015/01/riscv-testing-frameworks-bootcamp-jan2015.pdf
But lacks test format definition.

Test format (reverse engineered, cannot find docs):
- ELF file for each test *suite* to execute
- a `.dump` file with the matching assembly to understand the test
- entry-point seems to always be `0x80_00_00_00`
- test *case* is small part of it, and test number is contained in register `gp` during test
  - `1137` is used to indicate exceptions
- cases jump to `fail` if bad, with single-argument syscall calling convention to signal result:
  - `a0` is set to `(test_case_num << 1) | 1`
  - `a7` is set to the function ID, 93 (`exit`)
  - `ecall`, indicating test num and fail status
- if not to `fail`, we end at `pass`, again a syscall:
  - `a0` is set to `0`
  - `a7` is set to the function ID, 93 (`exit`)
  - `ecall`

The `.dump` files are the disassembled versions of the respective binary files; useful data during test debugging.



