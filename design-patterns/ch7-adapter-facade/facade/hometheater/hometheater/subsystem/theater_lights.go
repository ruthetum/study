package subsystem

import "fmt"

type TheaterLights struct {
}

func NewTheaterLights() *TheaterLights {
	return &TheaterLights{}
}

func (l *TheaterLights) On() {
	fmt.Println("Theater lights on")
}

func (l *TheaterLights) Dim(value int) {
	fmt.Println(fmt.Sprintf("Theater lights dim %v", value))
}
