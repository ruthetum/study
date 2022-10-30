package pizzafactory

import (
	"abstract-factory/cheese"
	"abstract-factory/dough"
	"abstract-factory/sauce"
	"fmt"
)

type IPizzaIngredientFactory interface {
	CreateDough() dough.IDough
	CreateSauce() sauce.ISauce
	CreateCheese() cheese.ICheese
}

func GetPizzaIngredientFactory(region string) (f IPizzaIngredientFactory, err error) {
	switch region {
	case "new york":
		f = NewNYPizzaIngredientFactory()
		return
	case "chicago":
		f = NewChicagoPizzaIngredientFactory()
		return
	}
	return nil, fmt.Errorf("invalid region")
}
