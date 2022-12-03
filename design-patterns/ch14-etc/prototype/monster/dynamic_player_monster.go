package monster

import "fmt"

type DynamicPlayerMonster struct {
	name string
}

func NewDynamicPlayerMonster(postfix string) DynamicPlayerMonster {
	return DynamicPlayerMonster{name: "dpm" + postfix}
}

func (m DynamicPlayerMonster) Print() {
	fmt.Println(m.name)
}

func (m DynamicPlayerMonster) Clone() Monster {
	return NewDynamicPlayerMonster("_clone")
}
