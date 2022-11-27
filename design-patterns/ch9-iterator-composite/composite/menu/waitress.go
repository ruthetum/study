package menu

type Waitress struct {
	menus Component
}

func NewWaitress(menus Component) *Waitress {
	return &Waitress{menus: menus}
}

func (w *Waitress) Print() {
	w.menus.Print()
}
