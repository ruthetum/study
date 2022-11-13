package subsystem

import "fmt"

type PopcornPopper struct {
}

func NewPopcornPopper() *PopcornPopper {
	return &PopcornPopper{}
}

func (p *PopcornPopper) On() {
	fmt.Println("Popcorn popper on")
}

func (p *PopcornPopper) Pop() {
	fmt.Println("Popcorn popper pop")
}

func (p *PopcornPopper) Off() {
	fmt.Println("Popcorn popper off")
}
