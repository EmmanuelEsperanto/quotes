package main

import (
	"log"
	"quotes/internal/app/apiserver"
)

func main() {
	err := apiserver.Start()
	if err != nil {
		log.Fatalf("Error stating apiserver: %v", err)
	}
	log.Println("API server started")
}
