package main

import (
	"hometheater/hometheater"
	"hometheater/hometheater/subsystem"
)

func main() {
	// subsystem
	amp := subsystem.NewAmplifier()
	tuner := subsystem.NewTuner()
	dvd := subsystem.NewDVDPlayer()
	cd := subsystem.NewCDPlayer()
	projector := subsystem.NewProjector()
	lights := subsystem.NewTheaterLights()
	screen := subsystem.NewScreen()
	popper := subsystem.NewPopcornPopper()

	// facade
	homeTheater := hometheater.NewHomeTheaterFacade(amp, tuner, dvd, cd, projector, lights, screen, popper)

	homeTheater.WatchMovie("Avengers")
	homeTheater.EndMovie()
}
