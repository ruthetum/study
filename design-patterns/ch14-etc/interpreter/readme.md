# Interpreter pattern
- 특정 언어의 문법 표현을 정의

## Structure
![image](https://upload.wikimedia.org/wikipedia/commons/thumb/b/bc/Interpreter_UML_class_diagram.svg/1072px-Interpreter_UML_class_diagram.svg.png)

## Example
```go
type ILogic interface {
    Evaluate() bool
}

var (
    Values = &V{
        vars: make(map[string]bool),
    }
)

type V struct {
    vars map[string]bool
}

func (v *V) Assign(key string, value bool) {
    v.vars[key] = value
}

func (v *V) Lookup(key string) bool {
    return v.vars[key]
}

type AndLogic struct {
	left  ILogic
	right ILogic
}

func NewAndLogic(left, right ILogic) AndLogic {
	return AndLogic{left: left, right: right}
}
func (l AndLogic) Evaluate() bool {
	return l.left.Evaluate() && l.right.Evaluate()
}

type Variable struct {
    name string
}

func NewVariable(name string) Variable {
    return Variable{name: name}
}

func (v Variable) ToString() string {
    return v.name
}

func (v Variable) Evaluate() bool {
    return Values.Lookup(v.name)
}

func main() {
    logic.Values.Assign("A", true)  // A = true
    logic.Values.Assign("B", false) // B = false

    // A && B
    term1 := logic.NewAndLogic(
        logic.NewVariable("A"),
        logic.NewVariable("B"),
    ).Evaluate()
    fmt.Println("A && B:", term1) // A && B: false

    ...
}
```

## Reference
- https://johngrib.github.io/wiki/pattern/interpreter/