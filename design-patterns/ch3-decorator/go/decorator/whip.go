package decorator

import "decorator/component"

type Whip struct {
	beverage component.Beverage
}

func AddWhip(b component.Beverage) *Whip {
	return &Whip{
		beverage: b,
	}
}

func (w *Whip) GetDescription() string {
	return w.beverage.GetDescription() + ", 휘핑"
}

func (w *Whip) Cost() float64 {
	cost := w.beverage.Cost()
	switch w.beverage.GetSize() {
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

func (w *Whip) GetSize() string {
	return w.beverage.GetSize()
}

func (w *Whip) SetSize(size int) {
	//TODO implement me
	panic("implement me")
}
