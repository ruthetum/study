package hometheater

import (
	"fmt"
	"hometheater/hometheater/subsystem"
)

type Facade struct {
	amp       *subsystem.Amplifier
	tuner     *subsystem.Tuner
	dvdPlayer *subsystem.DVDPlayer
	cdPlayer  *subsystem.CDPlayer
	projector *subsystem.Projector
	lights    *subsystem.TheaterLights
	screen    *subsystem.Screen
	popper    *subsystem.PopcornPopper
}

func NewHomeTheaterFacade(
	amp *subsystem.Amplifier,
	tuner *subsystem.Tuner,
	dvdPlayer *subsystem.DVDPlayer,
	cdPlayer *subsystem.CDPlayer,
	projector *subsystem.Projector,
	lights *subsystem.TheaterLights,
	screen *subsystem.Screen,
	popper *subsystem.PopcornPopper,
) *Facade {
	return &Facade{
		amp:       amp,
		tuner:     tuner,
		dvdPlayer: dvdPlayer,
		cdPlayer:  cdPlayer,
		projector: projector,
		lights:    lights,
		screen:    screen,
		popper:    popper,
	}
}

func (f *Facade) WatchMovie(movieName string) {
	fmt.Println("Get ready to watch a movie...")
	f.popper.On()
	f.popper.Pop()
	f.lights.Dim(10)
	f.screen.Down()
	f.projector.On()
	f.projector.WideScreenMode()
	f.amp.On()
	f.amp.SetDVD(f.dvdPlayer)
	f.amp.SetSurroundSound()
	f.amp.SetVolume(5)
	f.dvdPlayer.On()
	f.dvdPlayer.Play()
}

func (f *Facade) EndMovie() {
	fmt.Println("Shutting movie theater down...")
	f.popper.Off()
	f.lights.On()
	f.screen.Up()
	f.projector.Off()
	f.amp.Off()
	f.dvdPlayer.Stop()
	f.dvdPlayer.Eject()
	f.dvdPlayer.Off()
}
