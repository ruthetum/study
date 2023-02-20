package main

import (
	"log"
	"strconv"
	"stress/user"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var success uint64 = 0
	var wg sync.WaitGroup

	t := time.Now()
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			u := user.New(strconv.Itoa(id))
			if u.Start("room1") {
				atomic.AddUint64(&success, 1)
			}
		}(i)
	}
	wg.Wait()
	log.Println("success=", atomic.LoadUint64(&success), "time=", time.Since(t))

	//u := user.New(strconv.Itoa(1))
	//u.Start("room1")

}
