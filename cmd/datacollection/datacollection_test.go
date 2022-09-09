package datacollection

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"testing"
)

func TestDataCollection_streamMemPoolTxs(t *testing.T) {
	clientWrapper := NewClients()
	pendingTxs := clientWrapper.SubscribePendingTxHashes(100)
	for {
		fmt.Println("Channel length: ", len(pendingTxs))
		fmt.Println("Channel capacity: ", cap(pendingTxs))
		currentTxHash := <-pendingTxs
		fmt.Println(currentTxHash)
	}
}

// TODO This test runs forever due to for loop in SubscribeFullPendingTxs, rethink this
func TestDataCollection_SubscribeFullMemPoolTransactions(t *testing.T) {
	mempoolTxs := make(chan *types.Transaction, 20)
	SubscribeFullPendingTxs(mempoolTxs)
	for i := 1; i < 25; i++ {
		tx := <-mempoolTxs
		fmt.Println("------------------------------------------------------")
		fmt.Println(tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println(tx.Value().String())    // 10000000000000000
		fmt.Println(tx.Gas())               // 105000
		fmt.Println(tx.GasPrice().Uint64()) // 102000000000
		fmt.Println(tx.Nonce())             // 110644
		fmt.Println(tx.Data())              // []
		fmt.Println(tx.To().Hex())          // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e
	}

	fmt.Println(len(MissingTxs), " pending TXs not found on geth node:")
	for _, txHash := range MissingTxs {
		fmt.Println(txHash.Hex())
	}
}
