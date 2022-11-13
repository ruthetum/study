package memory

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type memory struct {
}

var memoryInstance *memory

func GetInstance() *memory {
	if memoryInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if memoryInstance == nil {
			fmt.Println("Creating single instance now.")
			memoryInstance = &memory{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return memoryInstance
}
