package menu

type PancakeHouseMenu struct {
	menus []*Menu
}

func NewPancakeHouseMenu() *PancakeHouseMenu {
	m := &PancakeHouseMenu{
		menus: make([]*Menu, 0),
	}

	m.AddItem("K&B 팬케이크 세트", "스크램블 에그와 토스트가 곁들여진 팬케이크", 2.99, true)
	m.AddItem("레귤러 팬케이크 세트", "달걀 후라이와 소시지가 곁들어 팬케이크", 2.99, false)
	m.AddItem("블루베리 팬케이크", "블루베리와 블루베리 시헙으로 만든 팬케이크", 3.49, true)
	m.AddItem("와플", "와플, 취향에 따라 블루베리나 딸기를 얹을 수 있습니다.", 3.59, true)

	return m
}

func (m *PancakeHouseMenu) CreateIterator() Iterator {
	if len(m.menus) == 0 {
		m.menus = make([]*Menu, 0)
	}
	return NewMenuIterator(m.menus)
}

func (m *PancakeHouseMenu) AddItem(name string, description string, price float64, vegetarian bool) {
	menu := NewMenu(name, description, price, vegetarian)
	m.menus = append(m.menus, menu)
}

func (m *PancakeHouseMenu) GetMenuItems() []*Menu {
	return m.menus
}
