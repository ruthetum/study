package menu

type Iterator interface {
	HasNext() bool
	GetNext() Component
}

type MenuIterator struct {
	index int
	menus []Component
}

func NewMenuIterator(menus []Component) *MenuIterator {
	return &MenuIterator{
		index: 0,
		menus: menus,
	}
}

func (m *MenuIterator) HasNext() bool {
	if m.index < len(m.menus) {
		if m.menus[m.index] != nil {
			return true
		}
	}
	return false

}
func (m *MenuIterator) GetNext() Component {
	if m.HasNext() {
		inMenu := m.menus[m.index]
		m.index++
		return inMenu
	}
	return nil
}
