package main

import (
	"compound/quackable/composite"
	"compound/quackable/decorator"
	"compound/quackable/duck/adpater"
	"compound/quackable/factory"
	"compound/quackable/goose"
	"compound/quackable/observer"
	"fmt"
)

func main() {
	// 오리 팩토리 생성
	countingDuckFactory := factory.CountingDuckFactory{}

	// 오리떼 생성
	mallardDuckOne := countingDuckFactory.CreateMallardDuck()
	mallardDuckTwo := countingDuckFactory.CreateMallardDuck()
	mallardDuckThree := countingDuckFactory.CreateMallardDuck()
	mallardDuckFour := countingDuckFactory.CreateMallardDuck()
	redheadDuck := countingDuckFactory.CreateRedheadDuck()
	duckCall := countingDuckFactory.CreateDuckCall()
	rubberDuck := countingDuckFactory.CreateRubberDuck()
	// 어뎁터 적용
	gooseDuck := adpater.GooseAdapter{Goose: goose.Goose{}}

	// 오리떼 집합 관리
	flockOfDucks := composite.NewFlock()
	flockOfDucks.Add(redheadDuck)
	flockOfDucks.Add(duckCall)
	flockOfDucks.Add(rubberDuck)
	flockOfDucks.Add(gooseDuck)

	flockOfMallards := composite.NewFlock()
	flockOfMallards.Add(mallardDuckOne)
	flockOfMallards.Add(mallardDuckTwo)
	flockOfMallards.Add(mallardDuckThree)
	flockOfMallards.Add(mallardDuckFour)

	flockOfDucks.Add(flockOfMallards)

	// 옵저버
	quackologist := &observer.Quackologist{}
	flockOfDucks.Register(quackologist)

	flockOfDucks.Quack()
	fmt.Println("Quack count:", decorator.NumberOfQuack)

	flockOfMallards.Quack()
	fmt.Println("Quack count:", decorator.NumberOfQuack)
}
