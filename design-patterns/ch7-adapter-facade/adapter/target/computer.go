package target

import "fmt"

type Computer interface {
	InsertIntoLightningPort()
}

type Mac struct {
}

func NewMac() *Mac {
	return &Mac{}
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}
