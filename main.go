package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade failed: ", err)
		return
	}
	defer conn.Close()

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read err: ", err)
			break
		}

		fmt.Printf("recieved: %s\n", err)
		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			fmt.Println("write err: ", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWS)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
