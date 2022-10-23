package behavior

import "fmt"

type QuackBehavior interface {
	Quack()
}

type Quack struct{}
type QuackMute struct{}
type Squeak struct{}

func (q Quack) Quack() {
	fmt.Println("꽥")
}

func (q QuackMute) Quack() {
	fmt.Println("<< 조용 >>")
}

func (q Squeak) Quack() {
	fmt.Println("삑")
}
