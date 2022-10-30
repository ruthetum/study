package cheese

type MozzarellaCheese struct {
	name string
}

func CreateMozzarellaCheese() ICheese {
	return &MozzarellaCheese{name: "모짜렐라 치즈"}
}

func (c MozzarellaCheese) ToString() string {
	return c.name
}
