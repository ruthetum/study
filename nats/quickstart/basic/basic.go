package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

func main() {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	// Simple Publisher
	msg := "Hello World"
	nc.Publish("foo", []byte(msg))
	fmt.Printf("Publish a message: %s\n", msg)

	// Simple Async Subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %v\n", m)
	})

	// Responding to a request message : 응답 기능, 요청에 대해서 편리하게 응답할 수 있음
	nc.Subscribe("request", func(m *nats.Msg) {
		m.Respond([]byte("answer is 42"))
	})

	// Simple Sync Subscriber : 지정한 subject에 대해서 구독, NextMsg()를 통해 동기적으로 수신
	sub, err := nc.SubscribeSync("foo")
	if err != nil {
		fmt.Printf("Error SubscribeSync(): %s\n", err)
	}
	// m, err := sub.NextMsg(timeout)
	m, err := sub.NextMsg(time.Second)
	fmt.Printf("Received a message: %v\n", m)

	sub.Unsubscribe()
	sub.Drain()

	// Channel Subscriber : 설정한 subject에 대해 구독을 설정하고 수신된 모든 메시지를 채널에 배치, Unsubscribe()가 호출될 때까지 채널을 닫으면 안 됨
	ch := make(chan *nats.Msg, 64)
	sub1, err := nc.ChanSubscribe("foo", ch)
	msg1 := <-ch
	fmt.Printf("Received a message: %s\n", string(msg1.Data))

	// Unsubscribe
	sub1.Unsubscribe()

	// Drain : 모든 관심사 제거, 메시지 처리될 때까지는 콜백 처리 유지
	sub1.Drain()

	// Requests
	msg2, err := nc.Request("help", []byte("help me"), 10*time.Millisecond)
	fmt.Printf("Request a message: %s\n", msg2.Data)

	// Replies
	nc.Subscribe("help", func(m *nats.Msg) {
		nc.Publish(m.Reply, []byte("I can help!"))
	})

	// Drain connection (Preferred for responders)
	// Close() not needed if this is called.
	nc.Drain()

	// Close connection : 계속해서 커넥션을 유지할 게 아니면 defer로 처리
	nc.Close()
}
