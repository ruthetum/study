package factory

import (
	"compound/quackable"
	"compound/quackable/decorator"
	"compound/quackable/duck"
	"compound/quackable/observer"
)

type CountingDuckFactory struct{}

func (f CountingDuckFactory) CreateMallardDuck() quackable.QuackAble {
	d := &duck.MallardDuck{}
	d.Observable = observer.NewObservable(d)
	return &decorator.QuackCounter{Duck: d}
}

func (f CountingDuckFactory) CreateRedheadDuck() quackable.QuackAble {
	d := &duck.RedheadDuck{}
	d.Observable = observer.NewObservable(d)
	return &decorator.QuackCounter{Duck: d}
}

func (f CountingDuckFactory) CreateDuckCall() quackable.QuackAble {
	d := &duck.DuckCall{}
	d.Observable = observer.NewObservable(d)
	return &decorator.QuackCounter{Duck: d}
}
func (f CountingDuckFactory) CreateRubberDuck() quackable.QuackAble {
	d := &duck.RedheadDuck{}
	d.Observable = observer.NewObservable(d)
	return &decorator.QuackCounter{Duck: d}
}
