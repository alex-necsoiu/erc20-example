package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	contract "unlocked/gen"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	ganacheUrl     = "http://localhost:7545"
	addr2          = "0xB976e687AE16528475B80896a6313799f8030f7c"
	ropstenAddress = "0xB976e687AE16528475B80896a6313799f8030f7c"
	contractAddres = "0xAF6925eb0ad743F9987fE5Fb0a39102cb13FF544"
	infuraUrl      = "https://ropsten.infura.io/v3/0c4ddf2996e14b63b79ad69cf5210281"

	pKey = "71a82eef9bdc4f237b57a4695fdcdb737dff4e96b8773c79108df1a9c6724074"
)

func main() {
	// Private Key
	// key, err := crypto.HexToECDSA(pKey)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Convert string into Address
	address := common.HexToAddress(ropstenAddress)
	cAddress := common.HexToAddress(contractAddres)

	// date := big.NewInt(1514764800)
	client, err := ethClient()
	if err != nil {
		log.Fatal(err)
	}

	instance, err := contract.NewUnlocked(cAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	balance, err := instance.BalanceOf(nil, address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balance Tx:", balance)

	// // Instantiate the contract and display its name
	// token, err := NewToken(common.HexToAddress("0x21e6fc92f93c8a1bb41e2be64b4e1f88a54d3576"), conn)
	// if err != nil {
	// 	log.Fatalf("Failed to instantiate a Token contract: %v", err)
	// }
	// name, err := token.Name(nil)
	// if err != nil {
	// 	log.Fatalf("Failed to retrieve token name: %v", err)
	// }
	// fmt.Println("Token name:", name)

}

// TxOpts take a client and private key and returns *bind.TransactOpts
func TxOpts(ctx context.Context, client *ethclient.Client, privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	fmt.Println("\n============================")
	fmt.Println("New Tx Options!")
	fmt.Println("\n============================")

	add := common.HexToAddress(ropstenAddress)

	nonce, err := client.PendingNonceAt(context.Background(), add)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	// auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	auth.Signer = Signer

	// auth.GasPrice = gasPrice
	// auth.GasLimit = uint64(6000000)
	// auth.Nonce = big.NewInt(int64(nonce))

	tx := &bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: auth.GasLimit,
		Value:    big.NewInt(10),
		Nonce:    auth.Nonce,
		Context:  ctx,
		GasPrice: auth.GasPrice,
	}

	return tx, nil

}

// // GetAllRooms returns an array of []Room
// func GetAllRooms(instance *contract.Bookingsystem) ([]contract.BookRoomRoom, error) {
// 	rooms, err := instance.AllRoomsByDate(nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Println("\n============================")
// 	fmt.Println("All Rooms:", rooms)
// 	fmt.Println("\n============================")

// 	return rooms, nil
// }

// // GetRoom returns an Room by index
// func GetRoom(instance *contract.Bookingsystem, idx int64) (*struct{Name:string bigInt}, error) {
// 	room, err := instance.GetRoom(nil, big.NewInt(idx))
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Println("\nRoom:", room)

// 	return room, nil
// }

// // CountRoom returns an Room
// func CountRoom(instance *contract.Bookingsystem) (uint8, error) {
// 	count, err := instance.RoomsCount(nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("\n============================")
// 	fmt.Println("Total Rooms:", count)
// 	fmt.Println("\n============================")
// 	return count, nil
// }

func ethClient() (*ethclient.Client, error) {

	// Convert string into Address
	// cAddress := common.HexToAddress(contractAddres)

	// Create Ethereum Client
	client, err := ethclient.DialContext(context.Background(), infuraUrl)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// instance, err := contract.NewBookingsystem(cAddress, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return client, nil
}
