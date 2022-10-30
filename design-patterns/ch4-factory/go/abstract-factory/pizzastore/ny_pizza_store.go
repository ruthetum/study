package pizzastore

import (
	"abstract-factory/pizza"
	"abstract-factory/pizzafactory"
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
	pizzaIngredientFactory := pizzafactory.NewNYPizzaIngredientFactory()
	switch menu {
	case "cheese":
		p = pizza.CreateCheesePizza(pizzaIngredientFactory)
		p.SetName("뉴욕 치즈 피자")
		return
	case "veggie":
		p = pizza.CreateVeggiePizza(pizzaIngredientFactory)
		p.SetName("뉴욕 야채 피자")
		return
	}
	return p, fmt.Errorf("invalid menu")
}
