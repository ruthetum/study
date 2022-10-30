package component

import "fmt"

type Decaf struct {
	description string
	size        int
}

func CreateDecaf() *Decaf {
	return &Decaf{
		description: "디카페인 커피",
		size:        0,
	}
}

func (b *Decaf) GetDescription() string {
	return b.description + fmt.Sprintf("[%v]", b.GetSize())
}

func (b *Decaf) Cost() float64 {
	return 1.05
}

func (b *Decaf) GetSize() string {
	return BeverageSizeMap[b.size]
}

func (b *Decaf) SetSize(size int) {
	b.size = size
}
