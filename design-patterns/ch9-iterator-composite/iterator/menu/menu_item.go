package menu

type Menu struct {
	name        string
	description string
	price       float64
	vegetarian  bool
}

func NewMenu(name string, description string, price float64, vegetarian bool) *Menu {
	return &Menu{
		name:        name,
		description: description,
		price:       price,
		vegetarian:  vegetarian,
	}
}

func (m *Menu) GetName() string {
	return m.name
}

func (m *Menu) GetDescription() string {
	return m.description
}

func (m *Menu) GetPrice() float64 {
	return m.price
}

func (m *Menu) IsVegetarian() bool {
	return m.vegetarian
}
