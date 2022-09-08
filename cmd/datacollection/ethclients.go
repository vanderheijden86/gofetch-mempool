package datacollection

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var infuraKey string

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	infuraKey = os.Getenv("INFURA_KEY")
}

func CreateEthClient() *ethclient.Client {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + infuraKey)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func CreateGethClient() *gethclient.Client {
	rpcClient, _ := rpc.Dial("wss://mainnet.infura.io/ws/v3/" + infuraKey)
	client := gethclient.New(rpcClient)
	return client
}
