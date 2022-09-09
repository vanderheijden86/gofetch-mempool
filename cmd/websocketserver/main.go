package main

import (
	"flag"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gorilla/websocket"
	"github.com/vanderheijden86/mempoolexplorer/cmd/datacollection"
	"log"
	"net/http"
)

// TODO put this in config file
var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var mempoolTxs = make(chan *types.Transaction, 20)

func sendMempoolTxs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()

	for {
		tx := <-mempoolTxs
		err = conn.WriteJSON(tx)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	go datacollection.SubscribeFullPendingTxs(mempoolTxs)
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/mempooltxs", sendMempoolTxs)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
