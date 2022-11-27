package adpater

import (
	"compound/quackable/goose"
	"compound/quackable/observer"
)

type GooseAdapter struct {
	Goose goose.Goose
}

func (a GooseAdapter) Quack() {
	a.Goose.Honk()
}

func (a GooseAdapter) Register(observer observer.Observer) {

}
