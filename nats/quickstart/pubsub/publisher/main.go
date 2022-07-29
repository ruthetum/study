package main

import (
	"github.com/nats-io/nats.go"
	"log"
)

func main() {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)
	if nc != nil {
		log.Println("Connected to " + nats.DefaultURL)
	}
	defer nc.Close()

	// Simple Publisher
	err := nc.Publish("foo", []byte("Hello World"))
	if err == nil {
		log.Println("Message published")
	}
}
