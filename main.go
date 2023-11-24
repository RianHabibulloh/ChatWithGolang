package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Message Received: ", string(msg))

		err = conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
	}
}