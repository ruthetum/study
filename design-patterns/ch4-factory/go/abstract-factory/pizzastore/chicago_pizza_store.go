package pizzastore

import (
	"abstract-factory/pizza"
	"abstract-factory/pizzafactory"
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
	pizzaIngredientFactory := pizzafactory.NewChicagoPizzaIngredientFactory()
	switch menu {
	case "cheese":
		p = pizza.CreateCheesePizza(pizzaIngredientFactory)
		p.SetName("시카코 치즈 피자")
		return
	case "veggie":
		p = pizza.CreateVeggiePizza(pizzaIngredientFactory)
		p.SetName("시카고 야채 피자")
		return
	}
	return p, fmt.Errorf("invalid menu")
}
