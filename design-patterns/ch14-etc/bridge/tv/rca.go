package tv

import "fmt"

type RCA struct {
}

func NewRCA() RCA {
	return RCA{}
}

func (tv RCA) On() {
	fmt.Println("RCA TV on")
}

func (tv RCA) Off() {
	fmt.Println("RCA TV off")
}

func (tv RCA) TuneChannel() {
	fmt.Println("RCA TV tune channel")
}
