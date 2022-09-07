package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var infuraKey string

func main() {
	godotenv.Load()
	infuraKey = os.Getenv("INFURA_KEY")
	pendingTxs := streamMemPoolTxs(createGethClient())
	// Wait 5 seconds to let the channel fill up and see how much TXs are in there.
	// TODO (2): Translate this into test cases
	// time.Sleep(5 * time.Second)

	for {
		fmt.Println("Channel length: ", len(pendingTxs))
		fmt.Println("Channel capacity: ", cap(pendingTxs))
		fmt.Println(<-pendingTxs)
	}
}

func createEthClient() *ethclient.Client {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + infuraKey)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func createGethClient() *gethclient.Client {
	rpcClient, _ := rpc.Dial("wss://mainnet.infura.io/ws/v3/" + infuraKey)
	client := gethclient.New(rpcClient)
	return client
}

// streamMemPoolTxs listens to all pending TXs that underlying ethereum node receives from incoming RPC requests and other nodes
func streamMemPoolTxs(geth *gethclient.Client) chan common.Hash {
	pendingTxs := make(chan common.Hash, 5_000)
	_, err := geth.SubscribePendingTransactions(context.Background(), pendingTxs)
	if err != nil {
		log.Fatal(err)
	}
	return pendingTxs
}
