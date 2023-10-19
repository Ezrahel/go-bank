package main

import (
	"log"
)

func main() {
	store, err := NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":8080", store)
	server.Run()

}
