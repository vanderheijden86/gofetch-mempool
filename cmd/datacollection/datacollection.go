package datacollection

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"log"
)

var MissingTxs = make([]common.Hash, 0, 20)
var Txs = make(map[common.Hash]*types.Transaction)

// StreamMemPoolTxHashes listens to all pending TXs that underlying ethereum node receives from incoming RPC requests and other nodes
func StreamMemPoolTxHashes(geth *gethclient.Client, bufferLength int) chan common.Hash {
	pendingTxs := make(chan common.Hash, bufferLength)
	_, err := geth.SubscribePendingTransactions(context.Background(), pendingTxs)
	if err != nil {
		log.Fatal(err)
	}
	return pendingTxs
}

func StoreTxDetails(txHash common.Hash) {
	client := CreateEthClient()
	tx, _, err := client.TransactionByHash(context.Background(), txHash)

	// If TX is not found, print the error, store the hash and return so the rest of the program can continue.
	if err == ethereum.NotFound {
		log.Println(err)
		storeMissingTxHashes(txHash)
		return
	} else if err != nil {
		log.Fatal(err)
	}
	Txs[txHash] = tx
}

func storeMissingTxHashes(txHash common.Hash) {
	MissingTxs = append(MissingTxs, txHash)
}
