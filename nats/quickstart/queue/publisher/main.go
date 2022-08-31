package main

import (
	"log"
	"strconv"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)
	if nc != nil {
		log.Println("Connected to " + nats.DefaultURL)
	}
	defer nc.Close()

	// Simple Publisher
	msg := "Hello World "
	for i := 1; i <= 35; i++ {
		index := strconv.Itoa(i)
		err := nc.Publish("foo", []byte(msg+index))
		if err == nil {
			log.Println("Message published " + index)
		}
	}
}
