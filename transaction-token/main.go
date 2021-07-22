package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	_ = client
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")

	privateKey, err := crypto.HexToECDSA("f5429bc3308335912ba05988e779063e53386bd70f27454a5c88aa6d12db3270")

	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1_000_000_000_000_000_000) // in wei -> 1 ETH
	gasLimit := uint64(21000)

	// gasprice := big.NewInt(30_000_000_000)  in wei -> 30 gwei
	gasprice, err := client.SuggestGasPrice(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x4f9111648a495984280e886D88FB6f010F682901")
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasprice, nil)

	chainID, err := client.NetworkID(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)

	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())
}
