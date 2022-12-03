package tv

import "fmt"

type Sony struct {
}

func NewSony() Sony {
	return Sony{}
}

func (tv Sony) On() {
	fmt.Println("Sony TV on")
}

func (tv Sony) Off() {
	fmt.Println("Sony TV off")
}

func (tv Sony) TuneChannel() {
	fmt.Println("Sony TV tune channel")
}
