package observer

import "fmt"

type Observer interface {
	Update(duck QuackObservable)
}

type Quackologist struct{}

func (l *Quackologist) Update(duck QuackObservable) {
	fmt.Println("Quackologist:", duck.GetName(), "just quacked")
}
