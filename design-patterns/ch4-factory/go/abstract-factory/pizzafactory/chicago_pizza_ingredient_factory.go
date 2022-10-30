package pizzafactory

import (
	"abstract-factory/cheese"
	"abstract-factory/dough"
	"abstract-factory/sauce"
)

type ChicagoPizzaIngredientFactory struct {
}

func NewChicagoPizzaIngredientFactory() IPizzaIngredientFactory {
	return &ChicagoPizzaIngredientFactory{}
}

func (f *ChicagoPizzaIngredientFactory) CreateDough() dough.IDough {
	return dough.CreateThickDough()
}

func (f *ChicagoPizzaIngredientFactory) CreateSauce() sauce.ISauce {
	return sauce.CreatePlumTomatoSauce()
}

func (f *ChicagoPizzaIngredientFactory) CreateCheese() cheese.ICheese {
	return cheese.CreateMozzarellaCheese()
}
