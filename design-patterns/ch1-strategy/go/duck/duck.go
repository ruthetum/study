package duck

import (
	"fmt"
	"strategy/behavior"
)

type Duck struct {
	Name          string
	flyBehavior   behavior.FlyBehavior
	quackBehavior behavior.QuackBehavior
}

func New(name string, flyBehavior behavior.FlyBehavior, quackBehavior behavior.QuackBehavior) *Duck {
	return &Duck{
		Name:          name,
		flyBehavior:   flyBehavior,
		quackBehavior: quackBehavior,
	}
}

func (d *Duck) Display() {
	fmt.Println("이름:", d.Name)
}

func (d *Duck) Swim() {
	fmt.Println("모든 오리는 물에 뜹니다.")
}

func (d *Duck) Fly() {
	d.flyBehavior.Fly()
}

func (d *Duck) Quack() {
	d.quackBehavior.Quack()
}

func (d *Duck) SetFlyBehavior(b behavior.FlyBehavior) {
	d.flyBehavior = b
}

func (d *Duck) SetQuackBehavior(b behavior.QuackBehavior) {
	d.quackBehavior = b
}
