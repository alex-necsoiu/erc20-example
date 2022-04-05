solc:
	@echo "Compiling Smart contracts and generating .BIN and .ABI files!"
	solc --bin --abi --overwrite contracts/Token.sol -o build
abigen:
	@echo "Generating Go files from ABI and BIN files!"
	abigen --pkg=unlocked --bin=build/Unlockd.bin --abi=build/Unlockd.abi  --out=gen/token_abi.go
testnet:
	go run ./deploy/$(ls -1 *.go | grep -v _tests.go)
run:	
	go run ./client/$(ls -1 *.go | grep -v _test.go)
test-client:
	go test -v ./client/*_test.go		