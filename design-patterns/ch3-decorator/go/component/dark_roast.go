package component

import "fmt"

type DarkRoast struct {
	description string
	size        int
}

func CreateDarkRoast() *DarkRoast {
	return &DarkRoast{
		description: "다크 로스트 커피",
		size:        0,
	}
}

func (b *DarkRoast) GetDescription() string {
	return b.description + fmt.Sprintf("[%v]", b.GetSize())
}

func (b *DarkRoast) Cost() float64 {
	return 0.99
}

func (b *DarkRoast) GetSize() string {
	return BeverageSizeMap[b.size]
}

func (b *DarkRoast) SetSize(size int) {
	b.size = size
}
