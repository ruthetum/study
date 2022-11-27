package menu

type Iterator interface {
	HasNext() bool
	GetNext() *Menu
}
