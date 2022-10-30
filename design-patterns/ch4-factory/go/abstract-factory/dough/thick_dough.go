package dough

type ThickDough struct {
	name string
}

func CreateThickDough() IDough {
	return &ThickDough{name: "두꺼운 도우"}
}

func (t ThickDough) ToString() string {
	return t.name
}
