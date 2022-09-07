package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/vanderheijden86/mempoolexplorer/cmd/datacollection"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func sendMempoolTxs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()

	for _, tx := range datacollection.Txs {
		err = conn.WriteJSON(tx)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	pendingTxs := datacollection.StreamMemPoolTxHashes(datacollection.CreateGethClient(), 10)
	for i := 1; i < 25; i++ {
		currentTxHash := <-pendingTxs
		fmt.Println(currentTxHash)
		datacollection.StoreTxDetails(currentTxHash)
	}
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/mempooltxs", sendMempoolTxs)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
