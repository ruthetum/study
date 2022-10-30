package decorator

import "decorator/component"

type Soy struct {
	beverage component.Beverage
}

func AddSoy(b component.Beverage) *Soy {
	return &Soy{
		beverage: b,
	}
}

func (s *Soy) GetDescription() string {
	return s.beverage.GetDescription() + ", 두유"
}

func (s *Soy) Cost() float64 {
	cost := s.beverage.Cost()
	switch s.beverage.GetSize() {
	case component.BeverageSizeMap[component.TallSize]:
		cost += 0.15
		break
	case component.BeverageSizeMap[component.GrandeSize]:
		cost += 0.2
		break
	case component.BeverageSizeMap[component.VentiSize]:
		cost += 0.25
		break
	}
	return cost
}

func (s *Soy) GetSize() string {
	return s.beverage.GetSize()
}

func (s *Soy) SetSize(size int) {
	//TODO implement me
	panic("implement me")
}
