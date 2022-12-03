package logic

type OrLogic struct {
	left  ILogic
	right ILogic
}

func NewOrLogic(left, right ILogic) OrLogic {
	return OrLogic{left: left, right: right}
}
func (l OrLogic) Evaluate() bool {
	return l.left.Evaluate() || l.right.Evaluate()
}
