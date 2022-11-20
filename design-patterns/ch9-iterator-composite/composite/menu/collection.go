package menu

type Collection interface {
	CreateIterator() Iterator
}
