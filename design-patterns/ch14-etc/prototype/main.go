package main

import "prototype/monster"

func main() {
	manager := monster.NewManager()

	wellKnownMonster := monster.NewWellKnownMonster("")
	dynamicPlayerMonster := monster.NewDynamicPlayerMonster("")

	manager.Register("wkm", wellKnownMonster)
	manager.Register("dpm", dynamicPlayerMonster)

	wellKnownMonster.Print()
	wellKnownMonsterClone := manager.Create("wkm")
	wellKnownMonsterClone.Print()
}
