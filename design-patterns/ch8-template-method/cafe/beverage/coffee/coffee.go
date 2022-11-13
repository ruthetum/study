package coffee

import "fmt"

type Coffee struct {
}

func New() *Coffee {
	return &Coffee{}
}

func (c *Coffee) BoilWater() {
	fmt.Println("물 끓이는 중")
}

func (c *Coffee) Brew() {
	fmt.Println("커피 우리는 중")
}

func (c *Coffee) PourInCup() {
	fmt.Println("컵에 따르는 중")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("설탕, 우유 넣는 중")
}
