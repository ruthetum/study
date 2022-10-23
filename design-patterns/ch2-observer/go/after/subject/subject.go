package subject

import "observer/observer"

type Subject interface {
	RegisterObserver(observer observer.Observer)
	RemoveObserver(observer observer.Observer)
	NotifyObserver()
}
