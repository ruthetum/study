package target

import "fmt"

type Duck interface {
	Quack()
	Fly()
}

type MallardDuck struct {
}

func NewMallardDuck() *MallardDuck {
	return &MallardDuck{}
}

func (d *MallardDuck) Quack() {
	fmt.Println("Quack")
}

func (d *MallardDuck) Fly() {
	fmt.Println("I'm flying")
}
