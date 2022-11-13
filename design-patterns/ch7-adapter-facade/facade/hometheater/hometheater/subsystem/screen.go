package subsystem

import "fmt"

type Screen struct {
}

func NewScreen() *Screen {
	return &Screen{}
}

func (s *Screen) Up() {
	fmt.Println("Screen up")
}

func (s *Screen) Down() {
	fmt.Println("Screen down")
}
