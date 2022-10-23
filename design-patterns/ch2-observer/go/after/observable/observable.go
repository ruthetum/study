package observable

import "observer/observer"

type Observable interface {
	RegisterObserver(observer observer.Observer)
	RemoveObserver(observer observer.Observer)
	NotifyAll()
}
