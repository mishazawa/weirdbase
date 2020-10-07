package api

import (
	"os"
	"log"
	"fmt"
	"net/http"
)

var MessageBus chan Message

func Run () {
	MessageBus = make(chan Message)
	http.HandleFunc("/ws", UpgradeConnection)
	go handleMessage()
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}

func handleMessage () {
	for {
		message := <-MessageBus
		switch message.Type {
		case "Create":
			if err := CreateRecord(message.Data.(map[string]interface{})); err != nil {
				log.Println("Create data err", err)
			}
		case "Read":
			log.Println("Read data")
		}
	}
}
