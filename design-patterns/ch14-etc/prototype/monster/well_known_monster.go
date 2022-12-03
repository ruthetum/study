package monster

import "fmt"

type WellKnownMonster struct {
	name string
}

func NewWellKnownMonster(postfix string) WellKnownMonster {
	return WellKnownMonster{name: "wkm" + postfix}
}

func (m WellKnownMonster) Print() {
	fmt.Println(m.name)
}

func (m WellKnownMonster) Clone() Monster {
	return NewWellKnownMonster("_clone")
}
