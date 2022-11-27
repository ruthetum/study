package decorator

import (
	"compound/quackable"
	"compound/quackable/observer"
	"sync/atomic"
)

var NumberOfQuack uint64 = 0

type QuackCounter struct {
	Duck quackable.QuackAble
}

func (c *QuackCounter) Quack() {
	c.Duck.Quack()
	atomic.AddUint64(&NumberOfQuack, 1)
}

func (c *QuackCounter) Register(observer observer.Observer) {
	c.Duck.Register(observer)
}
