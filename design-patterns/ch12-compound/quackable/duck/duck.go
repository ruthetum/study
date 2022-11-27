package duck

import (
	"compound/quackable/observer"
	"fmt"
)

type MallardDuck struct {
	Observable observer.QuackObservable
}

func (d *MallardDuck) Quack() {
	fmt.Println("Quack")
	d.NotifyAll()
}

func (d *MallardDuck) Register(observer observer.Observer) {
	d.RegisterObserver(observer)
}

func (d *MallardDuck) RegisterObserver(observer observer.Observer) {
	d.Observable.RegisterObserver(observer)
}

func (d *MallardDuck) NotifyAll() {
	d.Observable.NotifyAll()
}

func (d *MallardDuck) GetName() string {
	return "MallardDuck"
}

type RedheadDuck struct {
	Observable observer.QuackObservable
}

func (d *RedheadDuck) Quack() {
	fmt.Println("Quack")
	d.NotifyAll()
}

func (d *RedheadDuck) Register(observer observer.Observer) {
	d.RegisterObserver(observer)
}

func (d *RedheadDuck) GetObservable() observer.QuackObservable {
	return d.Observable
}

func (d *RedheadDuck) RegisterObserver(observer observer.Observer) {
	d.Observable.RegisterObserver(observer)
}

func (d *RedheadDuck) NotifyAll() {
	d.Observable.NotifyAll()
}

func (d *RedheadDuck) GetName() string {
	return "RedheadDuck"
}

type DuckCall struct {
	Observable observer.QuackObservable
}

func (d *DuckCall) Quack() {
	fmt.Println("Kwak")
	d.NotifyAll()
}

func (d *DuckCall) Register(observer observer.Observer) {
	d.RegisterObserver(observer)
}

func (d *DuckCall) RegisterObserver(observer observer.Observer) {
	d.Observable.RegisterObserver(observer)
}

func (d *DuckCall) NotifyAll() {
	d.Observable.NotifyAll()
}

func (d *DuckCall) GetName() string {
	return "DuckCall"
}

type RubberDuck struct {
	Observable observer.QuackObservable
}

func (d *RubberDuck) Quack() {
	fmt.Println("Squack")
	d.NotifyAll()
}

func (d *RubberDuck) Register(observer observer.Observer) {
	d.RegisterObserver(observer)
}

func (d *RubberDuck) RegisterObserver(observer observer.Observer) {
	d.Observable.RegisterObserver(observer)
}

func (d *RubberDuck) NotifyAll() {
	d.Observable.NotifyAll()
}

func (d *RubberDuck) GetName() string {
	return "RubberDuck"
}
