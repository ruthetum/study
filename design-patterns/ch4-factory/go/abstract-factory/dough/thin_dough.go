package dough

type ThinDough struct {
	name string
}

func CreateThinDough() IDough {
	return &ThinDough{name: "씬 도우"}
}

func (t ThinDough) ToString() string {
	return t.name
}
