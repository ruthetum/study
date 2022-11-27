package machine

import (
	"fmt"
)

type HasQuarterState struct {
	machine *GumballMachine
}

func NewHasQuarterState(machine *GumballMachine) State {
	return HasQuarterState{machine: machine}
}

func (s HasQuarterState) InsertQuarter() {
	fmt.Println("You can't insert another quarter")
}

func (s HasQuarterState) EjectQuarter() {
	fmt.Println("Quarter returned")
	s.machine.SetState(s.machine.GetNoQuarterState())
}

func (s HasQuarterState) TurnCrank() {
	fmt.Println("You turned...")
	s.machine.SetState(s.machine.GetSoldState())
}

func (s HasQuarterState) Dispense() {
	fmt.Println("No gumball dispensed")
}

func (s HasQuarterState) ToString() string {
	return "HasQuarterState"
}
