package composite

import (
	"compound/quackable"
	"compound/quackable/observer"
)

type Flock struct {
	quackers []quackable.QuackAble
}

func NewFlock() *Flock {
	return &Flock{
		quackers: make([]quackable.QuackAble, 0),
	}
}

func (f *Flock) Add(q quackable.QuackAble) {
	f.quackers = append(f.quackers, q)
}

func (f *Flock) Quack() {
	// iterator 패턴 적용 가능
	for _, q := range f.quackers {
		q.Quack()
	}
}

func (f *Flock) Register(observer observer.Observer) {
	f.RegisterObserver(observer)
}

func (f *Flock) GetObservable() observer.QuackObservable {
	return nil
}

func (f *Flock) RegisterObserver(observer observer.Observer) {
	for _, q := range f.quackers {
		q.Register(observer)
	}
}

func (f *Flock) NotifyAll() {
}
