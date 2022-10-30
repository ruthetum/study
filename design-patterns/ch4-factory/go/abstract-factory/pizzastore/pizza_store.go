package pizzastore

import "abstract-factory/pizza"

type IPizzaStore interface {
	OrderPizza(menu string) (p pizza.IPizza, err error)
	CreatePizza(menu string) (p pizza.IPizza, err error)
}
