package monitor

import (
	"fmt"
	"proxy/machine"
)

type GumballMonitor struct {
	machine *machine.GumballMachine
}

func NewGumballMonitor(machine *machine.GumballMachine) GumballMonitor {
	return GumballMonitor{
		machine: machine,
	}
}

func (m GumballMonitor) Report() {
	fmt.Println("\n[Report]")
	fmt.Println("뽑기 기계 위치:", m.machine.GetLocation())
	fmt.Println("현재 재고:", m.machine.GetCount())
	fmt.Println("현재 상태:", m.machine.GetState().ToString())
	fmt.Println()
}
