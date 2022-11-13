package subsystem

import "fmt"

type Projector struct {
}

func NewProjector() *Projector {
	return &Projector{}
}

func (p *Projector) On() {
	fmt.Println("Projector on")
}

func (p *Projector) WideScreenMode() {
	fmt.Println("Projector wide screen mode")
}

func (p *Projector) Off() {
	fmt.Println("Projector off")
}
