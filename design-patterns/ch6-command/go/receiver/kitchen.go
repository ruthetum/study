package receiver

import "fmt"

type Kitchen struct {
}

func NewKitchen() *Kitchen {
	return &Kitchen{}
}

func (t *Kitchen) On() {
	fmt.Println("Turning on Kitchen")
}

func (t *Kitchen) Off() {
	fmt.Println("Turning off Kitchen")
}
