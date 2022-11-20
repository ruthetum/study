package main

import (
	"fmt"
	"iterator/menu"
)

func main() {

	pancakeHouseMenu := menu.NewPancakeHouseMenu()
	pancakeIterator := pancakeHouseMenu.CreateIterator()
	for pancakeIterator.HasNext() {
		m := pancakeIterator.GetNext()
		fmt.Printf("Menu is %+v\n", m)
	}

	fmt.Println("\n======\n")

	dinerMenu := menu.NewDinerMenu()
	dinerIterator := dinerMenu.CreateIterator()
	for dinerIterator.HasNext() {
		m := dinerIterator.GetNext()
		fmt.Printf("Menu is %+v\n", m)
	}
}
