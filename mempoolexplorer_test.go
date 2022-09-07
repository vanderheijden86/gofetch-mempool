package mempoolexplorer

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestMain_streamMemPoolTxs(t *testing.T) {
	pendingTxs := streamMemPoolTxs(createGethClient(), 100)
	for {
		fmt.Println("Channel length: ", len(pendingTxs))
		fmt.Println("Channel capacity: ", cap(pendingTxs))
		currentTxHash := <-pendingTxs
		fmt.Println(currentTxHash)
	}
}

func TestMain_getTxDetails(t *testing.T) {
	// 0xaf745220755919ee3386ca28cc207e87388841832ee4bd67d7260b06b914af85
	printTxDetails(common.HexToHash("0xaf745220755919ee3386ca28cc207e87388841832ee4bd67d7260b06b914af85"))
}

func TestMain_getTxDetails_Live(t *testing.T) {
	pendingTxs := streamMemPoolTxs(createGethClient(), 10)
	for i := 1; i < 25; i++ {
		currentTxHash := <-pendingTxs
		fmt.Println(currentTxHash)
		printTxDetails(currentTxHash)
	}

	fmt.Println("Pending TXs not found on geth node:")
	for _, txHash := range missingTxs {
		fmt.Println(txHash.Hex())
	}

}
