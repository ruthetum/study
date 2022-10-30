package decorator

import (
	"decorator/component"
)

type Mocha struct {
	beverage component.Beverage
}

func AddMocha(b component.Beverage) *Mocha {
	return &Mocha{
		beverage: b,
	}
}

func (m *Mocha) GetDescription() string {
	return m.beverage.GetDescription() + ", 모카"
}

func (m *Mocha) Cost() float64 {
	cost := m.beverage.Cost()
	switch m.beverage.GetSize() {
	case component.BeverageSizeMap[component.TallSize]:
		cost += 0.2
		break
	case component.BeverageSizeMap[component.GrandeSize]:
		cost += 0.25
		break
	case component.BeverageSizeMap[component.VentiSize]:
		cost += 0.3
		break
	}
	return cost
}

func (m *Mocha) GetSize() string {
	return m.beverage.GetSize()
}

func (m *Mocha) SetSize(size int) {
	//TODO implement me
	panic("implement me")
}
