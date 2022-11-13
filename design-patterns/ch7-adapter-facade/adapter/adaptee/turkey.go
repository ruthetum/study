package adaptee

import "fmt"

type Turkey interface {
	Gobble()
	Fly()
}

type WildTurkey struct {
}

func NewWildTurkey() *WildTurkey {
	return &WildTurkey{}
}

func (t *WildTurkey) Gobble() {
	fmt.Println("Gobble gobble")
}

func (t *WildTurkey) Fly() {
	fmt.Println("I'm flying a short distance")
}
