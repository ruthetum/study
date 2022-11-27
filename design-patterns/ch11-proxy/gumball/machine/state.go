package machine

type State interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
	Dispense()
	ToString() string
}
