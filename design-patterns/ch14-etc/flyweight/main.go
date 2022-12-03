package main

import "flyweight/tree"

func main() {
	factory := tree.GetTreeFactory()
	factory.AddTree(1, 2)
	factory.AddTree(5, 1)
	factory.AddTree(3, 3)
	factory.Display()
}
