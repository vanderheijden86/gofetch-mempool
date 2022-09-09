package datacollection

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
)

var MissingTxs = make([]common.Hash, 0, 20)

// SubscribePendingTxHashes subscribes to "newPendingTransactions" events published by the Geth node.
// Returns the hash for all transactions that are added to the pending state and are signed with a key that is available in the node.
func (cs *ClientWrapper) SubscribePendingTxHashes(bufferLength int) chan common.Hash {
	pendingTxHashes := make(chan common.Hash, bufferLength)
	_, err := cs.geth.SubscribePendingTransactions(context.Background(), pendingTxHashes)
	if err != nil {
		log.Fatal(err)
	}
	return pendingTxHashes
}

// getFullTx returns the transaction with the given txHash. If the transaction isn't found, it adds the txHash to MissingTxs.
func (cs *ClientWrapper) getFullTx(txHash common.Hash) *types.Transaction {

	tx, _, err := cs.eth.TransactionByHash(context.Background(), txHash)

	// If TX is not found, print the error, store the hash and return so the rest of the program can continue.
	if err == ethereum.NotFound {
		log.Println(err)
		MissingTxs = append(MissingTxs, txHash)
		return nil
	} else if err != nil {
		log.Fatal(err)
	}
	return tx
}

// SubscribeFullPendingTxs subscribes to the nodes pendingTransactions, providing the full transactions (not just the hash).
func SubscribeFullPendingTxs(c chan *types.Transaction) {
	clientWrapper := NewClients()
	pendingTxHashes := clientWrapper.SubscribePendingTxHashes(10)
	for {
		currentTxHash := <-pendingTxHashes
		fmt.Println(currentTxHash)
		c <- clientWrapper.getFullTx(currentTxHash)
	}
}
