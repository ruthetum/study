package decorator

import "decorator/component"

type Milk struct {
	beverage component.Beverage
}

func AddMilk(b component.Beverage) *Milk {
	return &Milk{
		beverage: b,
	}
}

func (m *Milk) GetDescription() string {
	return m.beverage.GetDescription() + ", 우유"
}

func (m *Milk) Cost() float64 {
	cost := m.beverage.Cost()
	switch m.beverage.GetSize() {
	case component.BeverageSizeMap[component.TallSize]:
		cost += 0.1
		break
	case component.BeverageSizeMap[component.GrandeSize]:
		cost += 0.15
		break
	case component.BeverageSizeMap[component.VentiSize]:
		cost += 0.2
		break
	}
	return cost
}

func (m *Milk) GetSize() string {
	return m.beverage.GetSize()
}

func (m *Milk) SetSize(size int) {
	m.beverage.SetSize(size)
}
