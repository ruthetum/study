package pizzastore

import (
	"factory-method/pizza"
	"fmt"
)

type ChicagoPizzaStore struct {
}

func NewChicagoPizzaStore() IPizzaStore {
	return &ChicagoPizzaStore{}
}

func (s *ChicagoPizzaStore) OrderPizza(menu string) (p pizza.IPizza, err error) {
	p, err = s.CreatePizza(menu)
	p.Prepare()
	p.Bake()
	p.Cut()
	p.Box()
	return
}

func (s *ChicagoPizzaStore) CreatePizza(menu string) (p pizza.IPizza, err error) {
	switch menu {
	case "cheese":
		p = pizza.CreateChicagoCheesePizza()
		return
	case "veggie":
		p = pizza.CreateChicagoVeggiePizza()
		return
	}
	return p, fmt.Errorf("invalid menu")
}
