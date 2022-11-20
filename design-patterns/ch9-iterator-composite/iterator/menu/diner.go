package menu

import (
	"errors"
)

const dinerMenuMaxItems = 6

type DinerMenu struct {
	numberOfItems int
	menus         []*Menu
}

func NewDinerMenu() *DinerMenu {
	m := &DinerMenu{
		numberOfItems: 0,
		menus:         make([]*Menu, dinerMenuMaxItems),
	}

	_ = m.AddItem("채식주의자용 BLT", "통밀 위에 식물성 베이컨, 상추, 토마토를 얹은 메뉴", 2.99, true)
	_ = m.AddItem("BLT", "통밀 위에 베이컨, 상추, 토마토를 얹은 메뉴", 2.99, false)
	_ = m.AddItem("오늘의 스프", "감자 샐러드를 곁들인 오늘의 스프", 3.29, false)
	_ = m.AddItem("핫도그", "사워크라우트, 양념, 양파, 치즈가 곁들여진 핫도그", 3.05, false)

	return m
}

func (m *DinerMenu) CreateIterator() Iterator {
	if len(m.menus) == 0 {
		m.menus = make([]*Menu, 0)
	}
	return NewMenuIterator(m.menus)
}

func (m *DinerMenu) AddItem(name string, description string, price float64, vegetarian bool) error {
	menu := NewMenu(name, description, price, vegetarian)
	if m.numberOfItems >= dinerMenuMaxItems {
		return errors.New("죄송합니다. 메뉴가 꽉 찼습니다.")
	}
	m.menus[m.numberOfItems] = menu
	m.numberOfItems++
	return nil
}

func (m *DinerMenu) GetMenuItems() []*Menu {
	return m.menus
}
