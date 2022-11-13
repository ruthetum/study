package receiver

import "fmt"

type LivingRoom struct {
}

func NewLivingRoom() *LivingRoom {
	return &LivingRoom{}
}

func (t *LivingRoom) On() {
	fmt.Println("Turning on LivingRoom")
}

func (t *LivingRoom) Off() {
	fmt.Println("Turning off LivingRoom")
}
