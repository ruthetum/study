package pizza

import (
	"abstract-factory/pizzafactory"
	"fmt"
)

type NYCheesePizza struct {
	Pizza
	ingredientFactory pizzafactory.IPizzaIngredientFactory
}

func CreateCheesePizza(ingredientFactory pizzafactory.IPizzaIngredientFactory) IPizza {
	return &NYCheesePizza{
		Pizza:             Pizza{},
		ingredientFactory: ingredientFactory,
	}
}

func (p *NYCheesePizza) Prepare() {
	fmt.Println("준비 중", p.name)

	p.Pizza.dough = p.ingredientFactory.CreateDough().ToString()
	p.Pizza.sauce = p.ingredientFactory.CreateSauce().ToString()
	p.Pizza.cheese = p.ingredientFactory.CreateCheese().ToString()

	fmt.Println("도우 전달...")
	fmt.Println("소스 추가...")
	fmt.Println("치즈 추가...")
}
