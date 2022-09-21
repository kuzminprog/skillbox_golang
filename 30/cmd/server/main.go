package main

import (
	"30/internal/server"
	"log"
)

func main() {
	err := server.StartServer()

	if err != nil {
		log.Fatalln(err)
	}
}
