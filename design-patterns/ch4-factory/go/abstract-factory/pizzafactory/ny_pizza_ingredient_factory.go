package pizzafactory

import (
	"abstract-factory/cheese"
	"abstract-factory/dough"
	"abstract-factory/sauce"
)

type NYPizzaIngredientFactory struct {
}

func NewNYPizzaIngredientFactory() IPizzaIngredientFactory {
	return &NYPizzaIngredientFactory{}
}

func (f *NYPizzaIngredientFactory) CreateDough() dough.IDough {
	return dough.CreateThinDough()
}

func (f *NYPizzaIngredientFactory) CreateSauce() sauce.ISauce {
	return sauce.CreateMarinaraSauce()
}

func (f *NYPizzaIngredientFactory) CreateCheese() cheese.ICheese {
	return cheese.CreateParmesanCheese()
}
