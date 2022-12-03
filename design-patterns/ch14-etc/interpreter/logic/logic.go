package logic

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
