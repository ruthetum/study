package menu

import "fmt"

type Item struct {
	name        string
	description string
	price       float64
	vegetarian  bool
}

func NewItem(name string, description string, price float64, vegetarian bool) Component {
	return &Item{
		name:        name,
		description: description,
		price:       price,
		vegetarian:  vegetarian,
	}
}

func (i *Item) Add(c Component) {
	return
}

func (i *Item) Print() {
	fmt.Printf("Item is %+v\n", i)
}

func (i *Item) GetName() string {
	return i.name
}

func (i *Item) GetDescription() string {
	return i.description
}

func (i *Item) GetPrice() float64 {
	return i.price
}

func (i *Item) IsVegetarian() bool {
	return i.vegetarian
}
