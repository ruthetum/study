package main

type IMediator interface {
	Notify(string)
}

type Mediator struct {
	alarm     Alarm
	pot       CoffeePot
	calendar  Calendar
	sprinkler Sprinkler
}

func (m Mediator) Notify(name string) {
	if name == "alarm" {
		m.Alarm()
	}
}

func (m Mediator) Alarm() {
	m.calendar.CheckCalendar()
	m.sprinkler.CheckSprinkler()
	m.pot.StartCoffee()
}

type Alarm struct {
	mediator IMediator
}

func (a Alarm) OnEvent() {
	a.mediator.Notify("alarm")
}

type CoffeePot struct {
	mediator IMediator
}

func (p CoffeePot) StartCoffee() {

}

type Calendar struct {
	mediator IMediator
}

func (c Calendar) CheckCalendar() {

}

type Sprinkler struct {
	mediator IMediator
}

func (s Sprinkler) CheckSprinkler() {

}

func main() {

}
