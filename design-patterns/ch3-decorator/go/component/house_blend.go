package component

import "fmt"

type HouseBlend struct {
	description string
	size        int
}

func CreateHouseBlend() *HouseBlend {
	return &HouseBlend{
		description: "하우스 블렌드 커피",
		size:        0,
	}
}

func (b *HouseBlend) GetDescription() string {
	return b.description + fmt.Sprintf("[%v]", b.GetSize())
}

func (b *HouseBlend) Cost() float64 {
	return 0.89
}

func (b *HouseBlend) GetSize() string {
	return BeverageSizeMap[b.size]
}

func (b *HouseBlend) SetSize(size int) {
	b.size = size
}
