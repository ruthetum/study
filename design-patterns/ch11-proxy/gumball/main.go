package main

import (
	"fmt"
	"proxy/machine"
	"proxy/monitor"
)

func main() {
	gumballMachine := machine.NewGumballMachine("Seoul", 2)
	fmt.Printf("Gumball machine is %+v\n", gumballMachine)

	gumballMonitor := monitor.NewGumballMonitor(gumballMachine)

	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()

	gumballMonitor.Report()

	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()

	gumballMonitor.Report()
}
