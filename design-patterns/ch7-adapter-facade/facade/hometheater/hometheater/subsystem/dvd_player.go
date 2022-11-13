package subsystem

import "fmt"

type DVDPlayer struct {
}

func NewDVDPlayer() *DVDPlayer {
	return &DVDPlayer{}
}

func (d *DVDPlayer) On() {
	fmt.Println("DVD Player on")
}

func (d *DVDPlayer) Play() {
	fmt.Println("DVD Player play")
}

func (d *DVDPlayer) Stop() {
	fmt.Println("DVD Player stop")
}

func (d *DVDPlayer) Eject() {
	fmt.Println("DVD Player eject")
}

func (d *DVDPlayer) Off() {
	fmt.Println("DVD Player off")
}
