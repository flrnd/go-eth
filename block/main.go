package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")

	header, err := client.HeaderByNumber(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(header.Number.Int64())

	block, err := client.BlockByNumber(context.Background(), blockNumber)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Block")
	fmt.Printf("Number: %d\n", block.Number().Uint64())
	fmt.Printf("Time: %v\n", block.Time())
	fmt.Printf("Difficulty: %v\n", block.Difficulty().Uint64())
	fmt.Printf("Hash: %v\n", block.Hash().Hex())
	fmt.Printf("Total transactions: %v\n", len(block.Transactions()))
}
