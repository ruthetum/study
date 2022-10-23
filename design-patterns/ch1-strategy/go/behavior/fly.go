package behavior

import "fmt"

type FlyBehavior interface {
	Fly()
}

type FlyWithWings struct{}
type FlyNoWay struct{}
type FlyRocketPower struct{}

func (f FlyWithWings) Fly() {
	fmt.Println("날고 있어요!")
}

func (f FlyNoWay) Fly() {
	fmt.Println("저는 못 날아요.!")
}

func (f FlyRocketPower) Fly() {
	fmt.Println("로켓 속도롤 날아갑니다.")
}
