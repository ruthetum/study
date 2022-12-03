# Flyweight pattern
- 어떤 클래스의 인스턴스 한 개만 가지고 여러 개의 가상 인스턴스를 제공
- 크기가 작은 여러 개의 객체를 매번 생성하지 않고 최대한 공유하여 사용하도록 메모리 절약

## Structure
![image](https://refactoring.guru/images/patterns/diagrams/flyweight/structure-2x.png)

## Example
- 조경 애플리케이션을 만들 때 나무를 배치하는 경우 실제 나무들이 많이 심어지는 경우 메모리를 많이 차지
- 나무들의 역할은 딱히 없기 때문에 가상 나무 객체를 만들어서 좌표만 관리
```go
var instance Tree

type Tree struct {
    positionX int
    positionY int
    age       int
    status    string
    ...
}

func (t Tree) Display(x, y int) {
	fmt.Printf("Tree position x: %d, y: %d\n", x, y)
}

func GetTreeInstance() Tree {
	return instance
}

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


func main() {
    factory := tree.GetTreeFactory()
    factory.AddTree(1, 2)
    factory.AddTree(5, 1)
    factory.AddTree(3, 3)
    factory.Display()
}
```




## Reference
- https://refactoring.guru/ko/design-patterns/flyweight