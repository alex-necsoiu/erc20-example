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
	addr2          = "0xD890c3FC59FCBddf5Ce62aC9AFfC90DEbb7C88DE"
	rinkebyAddress = "0xAa3777F59260b8bD003e850E321AdBc576115b06"
	contractAddres = "0xAF6925eb0ad743F9987fE5Fb0a39102cb13FF544"
	infuraUrl      = "https://rinkeby.infura.io/v3/7db961138c114b7882db2ff9788cded0"

	pKey = "fd2179823adc4e772665848074ff54c3a0351a06accbb9e08498a51da1b519c5"
)

func main() {
	// Private Key
	// key, err := crypto.HexToECDSA(pKey)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Convert string into Address
	address := common.HexToAddress(rinkebyAddress)
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

	add := common.HexToAddress(rinkebyAddress)

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
