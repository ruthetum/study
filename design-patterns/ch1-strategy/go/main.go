package main

import (
	"strategy/behavior"
	"strategy/duck"
)

func main() {

	// Init mallard duck
	mallardDuckName := "mallard duck"
	flyWithWings := behavior.FlyWithWings{}
	quack := behavior.Quack{}
	mallard := duck.New(mallardDuckName, flyWithWings, quack)

	// Perform
	mallard.Display() // 이름: mallard duck
	mallard.Quack()   // 꽥
	mallard.Fly()     // 날고 있어요!

	// Init model duck
	modelDuckName := "model duck"
	flyNoWay := behavior.FlyNoWay{}
	squeak := behavior.Squeak{}
	model := duck.New(modelDuckName, flyNoWay, squeak)

	// Perform
	model.Display() // 이름: model duck
	model.Quack()   // 삑
	model.Fly()     // 저는 못 날아요.!

	// Update fly behavior
	model.SetFlyBehavior(behavior.FlyRocketPower{})
	model.Fly() // 로켓 속도롤 날아갑니다.

	// Update quack behavior
	model.SetQuackBehavior(behavior.QuackMute{})
	model.Quack() // << 조용 >>
}
