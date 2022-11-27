package main

import (
	"fmt"
	"sync/atomic"
)

type QuackAble interface {
	Quack()
}

type MallardDuck struct{}

func (d MallardDuck) Quack() {
	fmt.Println("Quack")
}

type RedheadDuck struct{}

func (d RedheadDuck) Quack() {
	fmt.Println("Quack")
}

type DuckCall struct{}

func (d DuckCall) Quack() {
	fmt.Println("Kwak")
}

type RubberDuck struct{}

func (d RubberDuck) Quack() {
	fmt.Println("Squack")
}

// Adapter
type Goose struct{}

func (g Goose) Honk() {
	fmt.Println("Honk")
}

type GooseAdapter struct {
	goose Goose
}

func (a GooseAdapter) Quack() {
	a.goose.Honk()
}

// Decorator
var numberOfQuack uint64 = 0

type QuackCounter struct {
	duck QuackAble
}

func (c *QuackCounter) Quack() {
	c.duck.Quack()
	atomic.AddUint64(&numberOfQuack, 1)
}

// Factory
type IDuckFactory interface {
	CreateMallardDuck() QuackAble
	CreateRedheadDuck() QuackAble
	CreateDuckCall() QuackAble
	CreateRubberDuck() QuackAble
}

type DuckFactory struct{}

func (f DuckFactory) CreateMallardDuck() QuackAble {
	return MallardDuck{}
}

func (f DuckFactory) CreateRedheadDuck() QuackAble {
	return RedheadDuck{}
}

func (f DuckFactory) CreateDuckCall() QuackAble {
	return DuckCall{}
}

func (f DuckFactory) CreateRubberDuck() QuackAble {
	return RedheadDuck{}
}

type CountingDuckFactory struct{}

func (f CountingDuckFactory) CreateMallardDuck() QuackAble {
	return &QuackCounter{MallardDuck{}}
}

func (f CountingDuckFactory) CreateRedheadDuck() QuackAble {
	return &QuackCounter{RedheadDuck{}}
}

func (f CountingDuckFactory) CreateDuckCall() QuackAble {
	return &QuackCounter{DuckCall{}}
}
func (f CountingDuckFactory) CreateRubberDuck() QuackAble {
	return &QuackCounter{RubberDuck{}}
}

// Composite
type Flock struct {
	quackers []QuackAble
}

func NewFlock() *Flock {
	return &Flock{
		quackers: make([]QuackAble, 0),
	}
}

func (f *Flock) Add(q QuackAble) {
	f.quackers = append(f.quackers, q)
}

func (f *Flock) Quack() {
	// iterator 패턴 적용 가능
	for _, q := range f.quackers {
		q.Quack()
	}
}

func main() {
	// QuackAble interface
	//mallardDuck := MallardDuck{}
	//redheadDuck := RedheadDuck{}
	//duckCall := DuckCall{}
	//rubberDuck := RubberDuck{}
	//gooseDuck := GooseAdapter{Goose{}}

	// Decorator
	//mallardDuck := &QuackCounter{MallardDuck{}}
	//redheadDuck := &QuackCounter{RedheadDuck{}}
	//duckCall := &QuackCounter{DuckCall{}}
	//rubberDuck := &QuackCounter{RubberDuck{}}
	//gooseDuck := GooseAdapter{Goose{}}

	// Factory
	//countingDuckFactory := CountingDuckFactory{}
	//mallardDuck := countingDuckFactory.CreateMallardDuck()
	//redheadDuck := countingDuckFactory.CreateRedheadDuck()
	//duckCall := countingDuckFactory.CreateDuckCall()
	//rubberDuck := countingDuckFactory.CreateRubberDuck()
	//gooseDuck := GooseAdapter{Goose{}}

	//mallardDuck.Quack()
	//redheadDuck.Quack()
	//duckCall.Quack()
	//rubberDuck.Quack()
	//gooseDuck.Quack()
	//fmt.Println("Quack count:", numberOfQuack)

	countingDuckFactory := CountingDuckFactory{}
	mallardDuckOne := countingDuckFactory.CreateMallardDuck()
	mallardDuckTwo := countingDuckFactory.CreateMallardDuck()
	mallardDuckThree := countingDuckFactory.CreateMallardDuck()
	mallardDuckFour := countingDuckFactory.CreateMallardDuck()
	redheadDuck := countingDuckFactory.CreateRedheadDuck()
	duckCall := countingDuckFactory.CreateDuckCall()
	rubberDuck := countingDuckFactory.CreateRubberDuck()
	gooseDuck := GooseAdapter{Goose{}}

	flockOfDucks := NewFlock()
	flockOfDucks.Add(redheadDuck)
	flockOfDucks.Add(duckCall)
	flockOfDucks.Add(rubberDuck)
	flockOfDucks.Add(gooseDuck)

	flockOfMallards := NewFlock()
	flockOfMallards.Add(mallardDuckOne)
	flockOfMallards.Add(mallardDuckTwo)
	flockOfMallards.Add(mallardDuckThree)
	flockOfMallards.Add(mallardDuckFour)

	flockOfDucks.Add(flockOfMallards)

	flockOfDucks.Quack()
	fmt.Println("Quack count:", numberOfQuack)

	flockOfMallards.Quack()
	fmt.Println("Quack count:", numberOfQuack)
}
