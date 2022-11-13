package tea

import "fmt"

type Tea struct {
}

func New() *Tea {
	return &Tea{}
}

func (t *Tea) BoilWater() {
	fmt.Println("물 끓이는 중")
}

func (t *Tea) Brew() {
	fmt.Println("차 우리는 중")
}

func (t *Tea) PourInCup() {
	fmt.Println("컵에 따르는 중")
}

func (t *Tea) AddCondiments() {
	fmt.Println("레몬 넣는 중")
}
