package main

import (
	"singleton/memory"
	"sync"
)

func main() {
	count := 10

	var wait sync.WaitGroup
	wait.Add(count)

	for i := 0; i < count; i++ {
		go func() {
			defer wait.Done()
			memory.GetInstance()
		}()

	}

	wait.Wait()
}
