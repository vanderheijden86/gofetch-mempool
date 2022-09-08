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

// StreamMemPoolTxHashes listens to all pending TXs that underlying ethereum node receives from incoming RPC requests and other nodes
func (cs *ClientWrapper) StreamMemPoolTxHashes(bufferLength int) chan common.Hash {
	pendingTxHashes := make(chan common.Hash, bufferLength)
	_, err := cs.geth.SubscribePendingTransactions(context.Background(), pendingTxHashes)
	if err != nil {
		log.Fatal(err)
	}
	return pendingTxHashes
}

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

func SubscribeFullMemPoolTransactions(c chan *types.Transaction) {
	clientWrapper := NewClients()
	pendingTxHashes := clientWrapper.StreamMemPoolTxHashes(10)
	for {
		currentTxHash := <-pendingTxHashes
		fmt.Println(currentTxHash)
		c <- clientWrapper.getFullTx(currentTxHash)
	}
}
