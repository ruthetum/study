package main

import (
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)
	if nc != nil {
		log.Println("Connected to " + nats.DefaultURL)
	}

	// Create a queue subscription on "foo" with queue name "workers"
	if _, err := nc.QueueSubscribe("foo", "workers", func(m *nats.Msg) {
		log.Println("Subscriber 1: ", string(m.Data))
	}); err != nil {
		log.Fatal(err)
	}

	// Create a queue subscription on "foo" with queue name "workers"
	if _, err := nc.QueueSubscribe("foo", "workers", func(m *nats.Msg) {
		log.Println("Subscriber 2: ", string(m.Data))
	}); err != nil {
		log.Fatal(err)
	}

	// Create a queue subscription on "foo" with queue name "workers"
	if _, err := nc.QueueSubscribe("foo", "workers", func(m *nats.Msg) {
		log.Println("Subscriber 3: ", string(m.Data))
	}); err != nil {
		log.Fatal(err)
	}

	// Keep the connection alive
	runtime.Goexit()
}
