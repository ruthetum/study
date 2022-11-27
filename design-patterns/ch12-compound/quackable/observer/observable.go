package observer

type QuackObservable interface {
	GetName() string
	RegisterObserver(observer Observer)
	NotifyAll()
}

type Observable struct {
	Observers []Observer
	Duck      QuackObservable
}

func NewObservable(d QuackObservable) *Observable {
	return &Observable{
		Observers: make([]Observer, 0),
		Duck:      d,
	}
}

func (o *Observable) GetName() string {
	return o.Duck.GetName()
}

func (o *Observable) RegisterObserver(observer Observer) {
	o.Observers = append(o.Observers, observer)
}

func (o *Observable) NotifyAll() {
	for _, observer := range o.Observers {
		observer.Update(o.Duck)
	}
}
