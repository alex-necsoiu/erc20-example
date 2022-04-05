# Build/Run

### Compiling Smart contracts and generating .BIN and .ABI files. Check the files generated in /build dir.

```shell
make solc
```
### Generate Go Bindings from .abi and .bin files.Check the files generated in /gen dir.

```shell
make abigen
```
### Deploy the Smart contract to desired environment.

```shell
make deploy
```
### Run a Client interaction with ERC20 contract.

```shell
make run
```
### Run the tests.

```shell
make Test
```
Check the Makefile for more details.
