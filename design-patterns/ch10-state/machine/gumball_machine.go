package machine

import (
	"fmt"
)

type GumballMachine struct {
	soldOutState    State
	noQuarterState  State
	hasQuarterState State
	soldState       State
	winnerState     State

	state State
	count int
}

func NewGumballMachine(count int) *GumballMachine {
	m := &GumballMachine{
		count: count,
	}

	m.soldOutState = NewSoldOutState(m)
	m.noQuarterState = NewNoQuarterState(m)
	m.hasQuarterState = NewHasQuarterState(m)
	m.soldState = NewSoldState(m)
	m.winnerState = NewWinnerState(m)

	if count > 0 {
		m.state = m.noQuarterState
	} else {
		m.state = m.soldOutState
	}
	return m
}

func (m *GumballMachine) InsertQuarter() {
	m.state.InsertQuarter()
}

func (m *GumballMachine) EjectQuarter() {
	m.state.EjectQuarter()
}

func (m *GumballMachine) TurnCrank() {
	m.state.TurnCrank()
	m.state.Dispense()
}
func (m *GumballMachine) ReleaseBall() {
	fmt.Println("a gumball comes rolling out the slot...")
	if m.count > 0 {
		m.count--
	}
}

func (m *GumballMachine) SetState(state State) {
	m.state = state
}

func (m *GumballMachine) GetCount() int {
	return m.count
}

func (m *GumballMachine) GetSoldOutState() State {
	return m.soldOutState
}

func (m *GumballMachine) GetNoQuarterState() State {
	return m.noQuarterState
}

func (m *GumballMachine) GetHasQuarterState() State {
	return m.hasQuarterState
}

func (m *GumballMachine) GetSoldState() State {
	return m.soldState
}

func (m *GumballMachine) GetWinnerState() State {
	return m.winnerState
}
