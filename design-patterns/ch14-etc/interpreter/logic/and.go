package logic

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
