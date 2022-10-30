package pizzastore

import (
	"factory-method/pizza"
	"fmt"
)

type NYPizzaStore struct {
}

func NewNYPizzaStore() IPizzaStore {
	return &NYPizzaStore{}
}

func (s *NYPizzaStore) OrderPizza(menu string) (p pizza.IPizza, err error) {
	p, err = s.CreatePizza(menu)
	p.Prepare()
	p.Bake()
	p.Cut()
	p.Box()
	return
}

func (s *NYPizzaStore) CreatePizza(menu string) (p pizza.IPizza, err error) {
	switch menu {
	case "cheese":
		p = pizza.CreateNYCheesePizza()
		return
	case "veggie":
		p = pizza.CreateNYVeggiePizza()
		return
	}
	return p, fmt.Errorf("invalid menu")
}
