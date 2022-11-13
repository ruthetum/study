package main

import (
	"adapter/adaptee"
	"adapter/adapter"
	"adapter/target"
	"fmt"
)

func main() {
	// Duck & Turkey begin
	description("duck & turkey example begin")
	mallardDuck := target.NewMallardDuck()

	wildTurkey := adaptee.NewWildTurkey()
	turkeyAdapter := adapter.NewTurkeyAdapter(wildTurkey)

	description("The turkey says")
	wildTurkey.Gobble()
	wildTurkey.Fly()

	description("The duck says")
	testDuck(mallardDuck)

	description("The turkey adapter says")
	testDuck(turkeyAdapter)

	description("duck & turkey example end")
	// Duck & Turkey end

	// Computer begin
	description("computer example begin")
	mac := target.NewMac()

	gram := adaptee.NewGram()
	windowsAdapter := adapter.NewWindowsAdapter(gram)

	description("mac")
	testComputer(mac)

	description("gram with windows adpater")
	testComputer(windowsAdapter)

	description("computer example end")
	// Computer end
}

func testDuck(duck target.Duck) {
	duck.Quack()
	duck.Fly()
}

func testComputer(computer target.Computer) {
	computer.InsertIntoLightningPort()
}

func description(msg string) {
	fmt.Println(fmt.Sprintf("\n[%v]", msg))
}
