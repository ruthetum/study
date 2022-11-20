package menu

type Component interface {
	Add(component Component)
	GetName() string
	GetDescription() string
	Print()
}
