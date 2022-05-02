package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Error starting Glimr!\nUsage:\nmake server\nmake client")
	}

	mode := os.Args[1]

	if mode == "client" {
		startClient()
	} else if mode == "server" {
		startServer()
	}
}
