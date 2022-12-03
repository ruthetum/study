package tree

import "fmt"

var instance Tree

type Tree struct {
	positionX int
	positionY int
	age       int
	status    string
}

func (t Tree) Display(x, y int) {
	fmt.Printf("Tree position x: %d, y: %d\n", x, y)
}

func GetTreeInstance() Tree {
	return instance
}
