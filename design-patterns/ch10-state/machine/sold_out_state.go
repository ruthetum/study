package machine

import (
	"fmt"
)

type SoldOutState struct {
	machine *GumballMachine
}

func NewSoldOutState(machine *GumballMachine) State {
	return SoldOutState{machine: machine}
}

func (s SoldOutState) InsertQuarter() {
	fmt.Println("You can't insert a quarter, the machine is sold out")
}

func (s SoldOutState) EjectQuarter() {
	fmt.Println("You can't eject, you haven't inserted a quarter yet")
}

func (s SoldOutState) TurnCrank() {
	fmt.Println("You turned, but there are no gumballs")
}

func (s SoldOutState) Dispense() {
	fmt.Println("No gumball dispensed")
}
