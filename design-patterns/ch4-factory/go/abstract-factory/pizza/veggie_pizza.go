package pizza

import (
	"abstract-factory/pizzafactory"
	"fmt"
)

type NYVeggiePizza struct {
	Pizza
	ingredientFactory pizzafactory.IPizzaIngredientFactory
}

func CreateVeggiePizza(ingredientFactory pizzafactory.IPizzaIngredientFactory) IPizza {
	return &NYVeggiePizza{
		Pizza:             Pizza{},
		ingredientFactory: ingredientFactory,
	}
}

func (p *NYVeggiePizza) Prepare() {
	fmt.Println("준비 중", p.name)

	p.Pizza.dough = p.ingredientFactory.CreateDough().ToString()
	p.Pizza.sauce = p.ingredientFactory.CreateSauce().ToString()
	p.Pizza.cheese = p.ingredientFactory.CreateCheese().ToString()

	fmt.Println("도우 전달...")
	fmt.Println("소스 추가...")
	fmt.Println("치즈 추가...")
}
