# Core NATS

## 1.  Publish - Subscribe
> https://docs.nats.io/nats-concepts/core-nats/pubsub

![img.png](https://683899388-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-LqMYcZML1bsXrN3Ezg0%2Fuploads%2Fgit-blob-22d59af386038cc2717176561ffc95c63c295926%2Fpubsub.svg?alt=media)

## Example

**publisher.go**

```go
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

```

**subscriber.go**

```go
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

```

## 2. Request-Reply
> https://docs.nats.io/nats-concepts/core-nats/reqreply

![](https://683899388-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-LqMYcZML1bsXrN3Ezg0%2Fuploads%2Fgit-blob-dc10798d4afca301adba55c1e85c599b25a2ae24%2Freqrepl.svg?alt=media)

- subscribe하고 있는 모든 서비스가 메시지는 전달받고, 그 중 하나의 서비스만 publisher에게 reply
- 동기식 요청(ex. HTTP 요청)에서 요청에 대한 응답이 필요한 경우 활용

### Example

**[basic]**

**request.go**

```go
package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	if nc != nil {
		log.Println("Connected to " + nats.DefaultURL)
	}
	defer nc.Close()

	// Send the request
	msg := "Hello World"
	reply, err := nc.Request("foo", []byte(msg), time.Second)
	if err != nil {
		log.Fatal(err)
	}

	// Use the response
	log.Printf("Reply: %s", reply.Data)
}
```

**reply.go**

```go
package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"strconv"

	"github.com/nats-io/nats.go"
)

func main() {
	id := strconv.Itoa(rand.Intn(100))
	log.Printf("[%s] start\n", id)
	nc, _ := nats.Connect(nats.DefaultURL)
	if nc != nil {
		log.Println("Connected to " + nats.DefaultURL)
	}

	count := 1
	_, err := nc.Subscribe("foo", func(msg *nats.Msg) {
		log.Println("Received msg:", string(msg.Data))

		respondErr := msg.Respond([]byte(id + " " + string(msg.Data) + strconv.Itoa(count)))
		if respondErr != nil {
			log.Println(respondErr)
		}
		count++
	})
	if err != nil {
		fmt.Println("error: ", err)
	}
	runtime.Goexit()
}
```

**[Queue group]**

**request.go**

```go
package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	if nc != nil {
		log.Println("Connected to " + nats.DefaultURL)
	}
	defer nc.Close()

	// Send the request
	msg := "Hello World"
	reply, err := nc.Request("foo", []byte(msg), time.Second)
	if err != nil {
		log.Fatal(err)
	}

	// Use the response
	log.Printf("Reply: %s", reply.Data)
}
```

**reply.go**

```go
package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"strconv"

	"github.com/nats-io/nats.go"
)

func main() {
	id := strconv.Itoa(rand.Intn(100))
	log.Printf("[%s] start\n", id)
	nc, _ := nats.Connect(nats.DefaultURL)
	if nc != nil {
		log.Println("Connected to " + nats.DefaultURL)
	}

	count := 1
	_, err := nc.QueueSubscribe("foo", "API", func(msg *nats.Msg) {
		log.Println("Received msg:", string(msg.Data))

		respondErr := msg.Respond([]byte(id + " " + string(msg.Data) + strconv.Itoa(count)))
		if respondErr != nil {
			log.Println(respondErr)
		}
		count++
	})
	if err != nil {
		fmt.Println("error: ", err)
	}
	runtime.Goexit()
}
```

## 3. Queue Groups
> https://docs.nats.io/nats-concepts/core-nats/queue

![](https://683899388-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-LqMYcZML1bsXrN3Ezg0%2Fuploads%2Fgit-blob-62652b3e6dd556e3cb1c3bb474ec10038334c600%2Fqueue.svg?alt=media)

- 한 서비스를 기준으로 여러 번의 처리를 진행할 필요가 없는 경우에 사용 (like consumer group in kafka)

### Example

**publisher.go**

```go
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
```

**subscriber.go**

```go
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

	// Create a queue subscription on "foo" with queue name "QAPI"
	if _, err := nc.QueueSubscribe("foo", "QAPI", func(m *nats.Msg) {
		log.Println("Subscriber 1: ", string(m.Data))
	}); err != nil {
		log.Fatal(err)
	}

	if _, err := nc.QueueSubscribe("foo", "QAPI", func(m *nats.Msg) {
		log.Println("Subscriber 2: ", string(m.Data))
	}); err != nil {
		log.Fatal(err)
	}

	if _, err := nc.QueueSubscribe("foo", "QAPI", func(m *nats.Msg) {
		log.Println("Subscriber 3: ", string(m.Data))
	}); err != nil {
		log.Fatal(err)
	}

	// Keep the connection alive
	runtime.Goexit()
}
```

## Message max size

*Messages have a maximum size (set in the server configuration with `max_payload`). The size is set to 1 MB by default but can be increased up to 64 MB if needed (though the NATS team recommends keeping the max message size to something more reasonable like 8 MB).*

### NATS Subscription

nats 는 subscribe 하고 있는 subject를 CS-trie 자료 구조로 들고 있어, 매번 메세지에 대해 redis pubsub 처럼 모든 subscribe를 순회할 필요가 없습니다. 결과적으로 N개의 client, M개의 subscriber 에 대해 O(N+M)의 시간복잡도가 아닌 O(log M) 의 복잡도를 갖게 됩니다

trie : https://www.geeksforgeeks.org/trie-insert-and-search/

## Reference

https://docs.nats.io/
https://github.com/nats-io/nats-server
https://github.com/nats-io/nats.go

### 설명

- 채널톡 실시간 채팅 서버 개선 여정 - 2편 : Nats.io로 Redis 대체하기 : https://channel.io/ko/blog/real-time-chat-server-2-redis-pub-sub
    - 채널톡은 결국 다시 redis로 (redis partitioning 적용)
- NATS 간략 설명 : [https://couplewith.tistory.com/entry/MSA-마이크로-서비스간의-빠른-메시징-처리를-위한-NATS](https://couplewith.tistory.com/entry/MSA-%EB%A7%88%EC%9D%B4%ED%81%AC%EB%A1%9C-%EC%84%9C%EB%B9%84%EC%8A%A4%EA%B0%84%EC%9D%98-%EB%B9%A0%EB%A5%B8-%EB%A9%94%EC%8B%9C%EC%A7%95-%EC%B2%98%EB%A6%AC%EB%A5%BC-%EC%9C%84%ED%95%9C-NATS)
    - MQ 설명 : https://acet.pe.kr/673
- 카프카와 비교 : [https://itnext.io/contrasting-nats-with-apache-kafka-1d3bdb9aa767#:~:text=TL%3BDR Kafka is an,strong ordering and persistence semantics](https://itnext.io/contrasting-nats-with-apache-kafka-1d3bdb9aa767#:~:text=TL%3BDR%20Kafka%20is%20an,strong%20ordering%20and%20persistence%20semantics).
- NATS Proxy : https://nats.io/blog/natsproxy_project/

### 예제

https://dev.to/karanpratapsingh/distributed-communication-patterns-with-nats-g17