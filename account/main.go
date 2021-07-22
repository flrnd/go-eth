package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func convert(b big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(b.String())
	return new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
}

func main() {
	client, err := ethclient.Dial("http://localhost:8545")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected")
	account := common.HexToAddress("0xaB798435FC3654010D133C10eee3d6e6D77d969C")
	blockNumber, _ := client.BlockNumber(context.Background())
	bigBlockNumber := new(big.Int).SetUint64(blockNumber)
	balance, err := client.BalanceAt(context.Background(), account, bigBlockNumber)

	if err != nil {
		log.Fatal(err)
	}

	ethValue := convert(*balance)

	fmt.Printf("Block number: %v\n", blockNumber)
	fmt.Printf("Account: %v\nBalance: %v ETH\n", account.Hex(), ethValue)

	pendingBalance, _ := client.PendingBalanceAt(context.Background(), account)

	fmt.Printf("Pending balance: %v ETH\n", convert(*pendingBalance))
}
