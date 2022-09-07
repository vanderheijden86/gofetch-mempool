package mempoolexplorer

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"log"
)

func main() {
	pendingTxs := streamMemPoolTxs(createGethClient(), 5)
	for {
		fmt.Println("Channel length: ", len(pendingTxs))
		fmt.Println("Channel capacity: ", cap(pendingTxs))
		fmt.Println(<-pendingTxs)
	}
}

// streamMemPoolTxs listens to all pending TXs that underlying ethereum node receives from incoming RPC requests and other nodes
func streamMemPoolTxs(geth *gethclient.Client, bufferLength int) chan common.Hash {
	pendingTxs := make(chan common.Hash, bufferLength)
	_, err := geth.SubscribePendingTransactions(context.Background(), pendingTxs)
	if err != nil {
		log.Fatal(err)
	}
	return pendingTxs
}

func printTxDetails(txHash common.Hash) {
	client := createEthClient()
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)

	// If TX is not found, just print the error and return so the rest of the program can continue. If there's an other error then log and exit.
	if err == ethereum.NotFound {
		log.Println(err)
		return
	} else if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tx.Hash().Hex())          // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	fmt.Println(tx.Value().String())      // 10000000000000000
	fmt.Println(tx.Gas())                 // 105000
	fmt.Println(tx.GasPrice().Uint64())   // 102000000000
	fmt.Println(tx.Nonce())               // 110644
	fmt.Println(tx.Data())                // []
	fmt.Println(tx.To().Hex())            // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e
	fmt.Println("isPending: ", isPending) // true
}
