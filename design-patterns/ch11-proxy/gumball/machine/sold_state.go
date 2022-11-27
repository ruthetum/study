package machine

import (
	"fmt"
)

type SoldState struct {
	machine *GumballMachine
}

func NewSoldState(machine *GumballMachine) State {
	return SoldState{machine: machine}
}

func (s SoldState) InsertQuarter() {
	fmt.Println("Please wait, we're already giving you a gumball")
}

func (s SoldState) EjectQuarter() {
	fmt.Println("Sorry, you already turned the crank")
}

func (s SoldState) TurnCrank() {
	fmt.Println("Turning twice doesn't get you another gumball!")
}

func (s SoldState) Dispense() {
	s.machine.ReleaseBall()
	if s.machine.GetCount() > 0 {
		s.machine.SetState(s.machine.GetNoQuarterState())
	} else {
		s.machine.SetState(s.machine.GetSoldOutState())
	}
}

func (s SoldState) ToString() string {
	return "SoldState"
}
