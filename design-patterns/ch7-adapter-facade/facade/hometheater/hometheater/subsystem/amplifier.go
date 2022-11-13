package subsystem

import "fmt"

type Amplifier struct {
}

func NewAmplifier() *Amplifier {
	return &Amplifier{}
}

func (a *Amplifier) On() {
	fmt.Println("Amplifier on")
}

func (a *Amplifier) SetDVD(dvd *DVDPlayer) {
	fmt.Println("Amplifier set DVD")
}

func (a *Amplifier) SetSurroundSound() {
	fmt.Println("Amplifier set surround sound")
}

func (a *Amplifier) SetVolume(value int) {
	fmt.Println(fmt.Sprintf("Amplifier set volume %v", value))
}

func (a *Amplifier) Off() {
	fmt.Println("Amplifier off")
}
