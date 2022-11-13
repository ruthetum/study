package adapter

import "adapter/adaptee"

type TurkeyAdapter struct {
	turkey adaptee.Turkey
}

func NewTurkeyAdapter(turkey adaptee.Turkey) *TurkeyAdapter {
	return &TurkeyAdapter{
		turkey: turkey,
	}
}

func (a *TurkeyAdapter) Quack() {
	a.turkey.Gobble()
}

func (a *TurkeyAdapter) Fly() {
	for i := 0; i < 5; i++ {
		a.turkey.Fly()
	}
}
