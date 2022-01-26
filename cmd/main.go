package main

import (
	"blockchain/server"
	"log"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}
