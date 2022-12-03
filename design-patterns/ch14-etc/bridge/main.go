package main

import (
	"bridge/controller"
	"bridge/tv"
)

func main() {

	rcaTV := tv.NewRCA()
	sonyTV := tv.NewSony()

	rcaRemoteController := controller.NewRCARemoteControl(rcaTV)
	rcaRemoteController.SetChannel()

	rcaRemoteController = controller.NewRCARemoteControl(sonyTV)
	rcaRemoteController.SetChannel()

	sonyRemoteController := controller.NewSonyRemoteControl(rcaTV)
	sonyRemoteController.SetChannel()

	sonyRemoteController = controller.NewSonyRemoteControl(sonyTV)
	sonyRemoteController.SetChannel()
}
