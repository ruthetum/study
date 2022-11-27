package factory

import (
	"compound/quackable"
	"compound/quackable/duck"
	"compound/quackable/observer"
)

type DuckFactory struct{}

func (f DuckFactory) CreateMallardDuck() quackable.QuackAble {
	d := &duck.MallardDuck{}
	d.Observable = observer.NewObservable(d)
	return d
}

func (f DuckFactory) CreateRedheadDuck() quackable.QuackAble {
	d := &duck.RedheadDuck{}
	d.Observable = observer.NewObservable(d)
	return d
}

func (f DuckFactory) CreateDuckCall() quackable.QuackAble {
	d := &duck.DuckCall{}
	d.Observable = observer.NewObservable(d)
	return d
}

func (f DuckFactory) CreateRubberDuck() quackable.QuackAble {
	d := &duck.RedheadDuck{}
	d.Observable = observer.NewObservable(d)
	return d
}
