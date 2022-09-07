# Ethereum Mempool Explorer

Golang implementation of an ethereum mempool explorer. For now it's mainly a way to gain experience with accessing 
the ethereum mempool (pending transactions), refresh my understanding of the go programming language
and work with the go-ethereum `ethclient` and `gethclient` packages.

## Setup

In the project directory, you need to place a file called `.env` in which you place your Infura key for ethereum 
mainnet. You can rename the `example.env` file to `.env` and paste in your own key.

```bash
INFURA_KEY=<place your INFURA key here>
```


## Running
You can either run the `mempoolexplorer.go` file to see the pending transactions as a stream of TX hashes, or run 
the tests in `mempoolexplorer_tests.go` to see the functionality to get transaction details. 