package tree

var (
	factoryInstance = &Factory{
		treeArray: make([][]int, 0),
	}
)

type Factory struct {
	treeArray [][]int
}

func GetTreeFactory() *Factory {
	return factoryInstance
}

func (f *Factory) AddTree(x, y int) {
	f.treeArray = append(f.treeArray, []int{x, y})
}

func (f *Factory) Display() {
	for _, t := range f.treeArray {
		GetTreeInstance().Display(t[0], t[1])
	}
}
