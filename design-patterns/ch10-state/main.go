package main

import (
	"fmt"
	"state/machine"
)

func main() {

	gumballMachine := machine.NewGumballMachine(10)
	fmt.Printf("Gumball machine is %+v\n", gumballMachine)

	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()

	fmt.Printf("Gumball machine is %+v\n", gumballMachine)

	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()

	fmt.Printf("Gumball machine is %+v\n", gumballMachine)

	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()

	fmt.Printf("Gumball machine is %+v\n", gumballMachine)

	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()

	fmt.Printf("Gumball machine is %+v\n", gumballMachine)

	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()

	fmt.Printf("Gumball machine is %+v\n", gumballMachine)
}
