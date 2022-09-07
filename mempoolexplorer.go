package mempoolexplorer

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"log"
)

var missingTxs = make([]common.Hash, 0, 20)
var txs = make(map[common.Hash]*types.Transaction)

// streamMemPoolTxHashes listens to all pending TXs that underlying ethereum node receives from incoming RPC requests and other nodes
func streamMemPoolTxHashes(geth *gethclient.Client, bufferLength int) chan common.Hash {
	pendingTxs := make(chan common.Hash, bufferLength)
	_, err := geth.SubscribePendingTransactions(context.Background(), pendingTxs)
	if err != nil {
		log.Fatal(err)
	}
	return pendingTxs
}

func storeTxDetails(txHash common.Hash) {
	client := createEthClient()
	tx, _, err := client.TransactionByHash(context.Background(), txHash)

	// If TX is not found, print the error, store the hash and return so the rest of the program can continue.
	if err == ethereum.NotFound {
		log.Println(err)
		storeMissingTxHashes(txHash)
		return
	} else if err != nil {
		log.Fatal(err)
	}
	txs[txHash] = tx
}

func storeMissingTxHashes(txHash common.Hash) {
	missingTxs = append(missingTxs, txHash)
}
