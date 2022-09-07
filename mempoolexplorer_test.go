package mempoolexplorer

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestMempoolExplorer_streamMemPoolTxs(t *testing.T) {
	pendingTxs := streamMemPoolTxHashes(createGethClient(), 100)
	for {
		fmt.Println("Channel length: ", len(pendingTxs))
		fmt.Println("Channel capacity: ", cap(pendingTxs))
		currentTxHash := <-pendingTxs
		fmt.Println(currentTxHash)
	}
}

func TestMempoolExplorer_storeTxDetails(t *testing.T) {
	// 0xaf745220755919ee3386ca28cc207e87388841832ee4bd67d7260b06b914af85
	storeTxDetails(common.HexToHash("0xaf745220755919ee3386ca28cc207e87388841832ee4bd67d7260b06b914af85"))
}

func TestMempoolExplorer_storeTxDetails_Live(t *testing.T) {
	pendingTxs := streamMemPoolTxHashes(createGethClient(), 10)
	for i := 1; i < 25; i++ {
		currentTxHash := <-pendingTxs
		fmt.Println(currentTxHash)
		storeTxDetails(currentTxHash)
	}

	fmt.Println(len(txs), " stored mempool TXs found on geth node:")
	for _, tx := range txs {
		fmt.Println("------------------------------------------------------")
		fmt.Println(tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println(tx.Value().String())    // 10000000000000000
		fmt.Println(tx.Gas())               // 105000
		fmt.Println(tx.GasPrice().Uint64()) // 102000000000
		fmt.Println(tx.Nonce())             // 110644
		fmt.Println(tx.Data())              // []
		fmt.Println(tx.To().Hex())          // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e
	}

	fmt.Println(len(missingTxs), " pending TXs not found on geth node:")
	for _, txHash := range missingTxs {
		fmt.Println(txHash.Hex())
	}
}
