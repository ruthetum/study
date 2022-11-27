package machine

import (
	"fmt"
)

type WinnerState struct {
	machine *GumballMachine
}

func NewWinnerState(machine *GumballMachine) State {
	return WinnerState{machine: machine}
}

func (s WinnerState) InsertQuarter() {
	fmt.Println("Please wait, we're already giving you a Gumball")
}

func (s WinnerState) EjectQuarter() {
	fmt.Println("Please wait, we're already giving you a Gumball")
}

func (s WinnerState) TurnCrank() {
	fmt.Println("Turning again doesn't get you another gumball!")
}

func (s WinnerState) Dispense() {
	fmt.Println("YOU'RE A WINNER! You get two gumballs for your quarter")
	s.machine.ReleaseBall()
	if s.machine.GetCount() == 0 {
		s.machine.SetState(s.machine.GetSoldOutState())
	} else {
		s.machine.ReleaseBall()
		if s.machine.GetCount() > 0 {
			s.machine.SetState(s.machine.GetNoQuarterState())
		} else {
			fmt.Println("Oops, out of gumballs!")
			s.machine.SetState(s.machine.GetSoldOutState())
		}
	}
}

func (s WinnerState) ToString() string {
	return "WinnerState"
}
