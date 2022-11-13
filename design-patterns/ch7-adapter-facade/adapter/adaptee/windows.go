package adaptee

import "fmt"

type Windows interface {
	InsertIntoUSBPort()
}

type Gram struct {
}

func NewGram() *Gram {
	return &Gram{}
}

func (w *Gram) InsertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}
