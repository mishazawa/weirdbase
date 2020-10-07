package main

import (
	"log"

	"github.com/mishazawa/weirdbase/api"
)

func main () {
	log.Println("Weirdbase")
	err := api.Connect()
	if err != nil {
		log.Println("Error connection: ", err)
		return
	} else {
		log.Println("Ping Ok.")
	}

	api.Run()
}
