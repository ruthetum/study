package quackable

import "compound/quackable/observer"

type QuackAble interface {
	Quack()
	Register(observer observer.Observer)
}
