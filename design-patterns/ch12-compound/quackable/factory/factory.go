package factory

import "compound/quackable"

// Factory
type IDuckFactory interface {
	CreateMallardDuck() quackable.QuackAble
	CreateRedheadDuck() quackable.QuackAble
	CreateDuckCall() quackable.QuackAble
	CreateRubberDuck() quackable.QuackAble
}
