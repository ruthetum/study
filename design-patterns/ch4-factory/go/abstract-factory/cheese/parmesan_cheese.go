package cheese

type ParmesanCheese struct {
	name string
}

func CreateParmesanCheese() ICheese {
	return &ParmesanCheese{name: "파마산 치즈"}
}

func (c ParmesanCheese) ToString() string {
	return c.name
}
