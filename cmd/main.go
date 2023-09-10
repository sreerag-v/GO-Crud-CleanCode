package main

import (
	di "Clean/Mongo-Crud/pkg/di"
	"log"
)

func main() {
	server, err := di.InitializeAPI()
	if err != nil {
		log.Fatal("error")
	}
	server.Start()
}
