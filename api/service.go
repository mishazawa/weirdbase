package api

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: checkCors,
}

func UpgradeConnection (w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	handleWebsocketConnection(conn)
}

func handleWebsocketConnection (conn *websocket.Conn) {
	log.Println("Connection was created.")

	done := make(chan interface{})

	go func (conn *websocket.Conn) {
		for {
			var m Message
			if err := conn.ReadJSON(&m); err != nil {
				log.Println(err)
				break
			}
			MessageBus <- m
		}
		close(done)
	}(conn)

	loop:
	for {
		select {
		case <- done:
			log.Println("Connection was closed.")
			break loop
		}
	}
}

func checkCors (r *http.Request) bool {
	return true
}
