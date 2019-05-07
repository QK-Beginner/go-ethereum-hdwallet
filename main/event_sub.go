package main

import (
	"context"
	_ "encoding/hex"
	"fmt"
	"log"


	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens string
}



func main() {
	//client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
	client, err := ethclient.Dial("wss://mainnet.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
			fmt.Printf("Log Index: %d\n", vLog.Index)
			fmt.Printf("Log Hash: %s\n", vLog.TxHash.Hex())

			switch vLog.Topics[0].Hex() {
			case logTransferSigHash.Hex():
				fmt.Printf("Log Name: Transfer\n")

				var transferEvent LogTransfer

				//err := contractAbi.Unpack(&transferEvent, "Transfer", vLog.Data)
				//if err != nil {
				//	log.Fatal(err)
				//}

				transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
				transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
				transferEvent.Tokens = common.Bytes2Hex(vLog.Data)

				fmt.Printf("From: %s\n", transferEvent.From.Hex())
				fmt.Printf("To: %s\n", transferEvent.To.Hex())
				fmt.Printf("Tokens: %s\n", transferEvent.Tokens)

				fmt.Printf("\n\n")

				//fmt.Println(time.Now(),vLog) // pointer to event log
			}
		}
	}
}
