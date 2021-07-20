package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func convert(b big.Int, e int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(b.String())
	return new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(e)))
}

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
	for _, tx := range block.Transactions() {
		fmt.Printf("transaction: %s\n", tx.Hash().Hex())
		fmt.Printf("transfered: %f ETH\n", convert(*tx.Value(), 18)) // wei / 10^18
		fmt.Printf("Gas used: %v\n", tx.Gas())
		fmt.Printf("Gas price: %v gwei\n", convert(*tx.GasPrice(), 9)) // gwei -> value / 10^9
		fmt.Printf("Nonce: %v\n", tx.Nonce())
		fmt.Printf("Data: %v\n", tx.Data())
		fmt.Printf("To: %s\n", tx.To().Hex())

		chainID, err := client.NetworkID(context.Background())

		if err != nil {
			log.Fatal(err)
		}

		if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), block.BaseFee()); err != nil {
			fmt.Println(msg.From().Hex())
		}
	}
}
