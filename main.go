package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/ws", handleWS)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
