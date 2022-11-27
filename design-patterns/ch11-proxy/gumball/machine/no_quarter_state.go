package machine

import (
	"fmt"
)

type NoQuarterState struct {
	machine *GumballMachine
}

func NewNoQuarterState(machine *GumballMachine) State {
	return NoQuarterState{machine: machine}
}

func (s NoQuarterState) InsertQuarter() {
	fmt.Println("You inserted a quarter")
	s.machine.SetState(s.machine.GetHasQuarterState())
}

func (s NoQuarterState) EjectQuarter() {
	fmt.Println("You haven't inserted a quarter")
}

func (s NoQuarterState) TurnCrank() {
	fmt.Println("You turned, but there's no quarter")
}

func (s NoQuarterState) Dispense() {
	fmt.Println("You need to pay first")
}

func (s NoQuarterState) ToString() string {
	return "NoQuarterState"
}
