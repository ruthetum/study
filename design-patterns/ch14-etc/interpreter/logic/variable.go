package logic

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
