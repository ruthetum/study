package menu

import "fmt"

type Menu struct {
	name        string
	description string
	sub         []Component
}

func NewMenu(name, description string) Component {
	return &Menu{
		name:        name,
		description: description,
		sub:         make([]Component, 0),
	}
}

func (m *Menu) GetName() string {
	return m.name
}

func (m *Menu) GetDescription() string {
	return m.description
}

func (m *Menu) Add(c Component) {
	m.sub = append(m.sub, c)
}

func (m *Menu) Print() {
	fmt.Printf("Menu is %+v\n", m)

	iterator := NewMenuIterator(m.sub)
	for iterator.HasNext() {
		s := iterator.GetNext()
		s.Print()
	}

}
