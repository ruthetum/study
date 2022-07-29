package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"runtime"
)

func main() {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)
	if nc != nil {
		log.Println("Connected to " + nats.DefaultURL)
	}

	// Simple Async Subscriber
	nc.Subscribe("foo", func(msg *nats.Msg) {
		log.Println(string(msg.Data))
	})

	// Keep the connection alive
	runtime.Goexit()
}
