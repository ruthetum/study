package component

import "fmt"

type Espresso struct {
	description string
	size        int
}

func CreateEspresso() *Espresso {
	return &Espresso{
		description: "에스프레소",
		size:        0,
	}
}

func (b *Espresso) GetDescription() string {
	return b.description + fmt.Sprintf("[%v]", b.GetSize())
}

func (b *Espresso) Cost() float64 {
	return 1.99
}

func (b *Espresso) GetSize() string {
	return BeverageSizeMap[b.size]
}

func (b *Espresso) SetSize(size int) {
	b.size = size
}
