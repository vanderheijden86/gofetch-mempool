package datacollection

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var ethNodeHttps string
var ethNodeWs string

type ClientWrapper struct {
	eth  *ethclient.Client
	geth *gethclient.Client
}

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	ethNodeHttps = os.Getenv("ETH_NODE_HTTPS")
	ethNodeWs = os.Getenv("ETH_NODE_WS")
}

func NewClients() *ClientWrapper {
	return &ClientWrapper{
		CreateEthClient(),
		CreateGethClient(),
	}
}

func CreateEthClient() *ethclient.Client {
	client, err := ethclient.Dial(ethNodeHttps)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func CreateGethClient() *gethclient.Client {
	rpcClient, _ := rpc.Dial(ethNodeWs)
	client := gethclient.New(rpcClient)
	return client
}
