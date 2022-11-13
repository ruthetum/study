package main

import (
	"cafe/beverage"
	"cafe/beverage/coffee"
	"cafe/beverage/tea"
)

func main() {
	coffee := beverage.New(coffee.New())
	tea := beverage.New(tea.New())

	coffee.PrepareRecipe()
	tea.PrepareRecipe()
}
