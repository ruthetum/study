package beverage

type ICaffeineBeverage interface {
	BoilWater()
	Brew()
	PourInCup()
	AddCondiments()
}

type CaffeineBeverage struct {
	beverage ICaffeineBeverage
}

func New(beverage ICaffeineBeverage) *CaffeineBeverage {
	return &CaffeineBeverage{
		beverage: beverage,
	}
}

func (b *CaffeineBeverage) PrepareRecipe() {
	b.beverage.BoilWater()
	b.beverage.Brew()
	b.beverage.PourInCup()
	b.beverage.AddCondiments()
}
