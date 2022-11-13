package main

import (
	"command/command"
	"command/invoker"
	"command/receiver"
)

func main() {
	// 리모컨 객체 생성 - invoker(호출자)
	remote := invoker.NewRemoteControl()

	// 커맨드 기능을 실행하기 위한 객체 생성 - receiver(수신자)
	livingRoom := receiver.NewLivingRoom()
	// 실제로 실행되는 기능을 구현하는 객체 생성 - command
	livingRoomLightOn := command.NewLightOnCommand(livingRoom)
	livingRoomLightOff := command.NewLightOffCommand(livingRoom)

	kitchen := receiver.NewKitchen()
	kitchenLightOn := command.NewLightOnCommand(kitchen)
	kitchenLightOff := command.NewLightOffCommand(kitchen)

	stereoPlayer := receiver.NewStereoPlayer()
	stereoPlayerOn := command.NewStereoOnWithCDCommand(stereoPlayer)
	stereoPlayerOff := command.NewStereoOffWithCDCommand(stereoPlayer)

	// 호출자(invoker)에 커맨드에 맵핑
	_ = remote.SetOnCommand(0, livingRoomLightOn)
	_ = remote.SetOffCommand(0, livingRoomLightOff)

	_ = remote.SetOnCommand(1, kitchenLightOn)
	_ = remote.SetOffCommand(1, kitchenLightOff)

	_ = remote.SetOnCommand(2, stereoPlayerOn)
	_ = remote.SetOffCommand(2, stereoPlayerOff)

	// 호출자(invoker)에서 원하는 커맨드 실행
	_ = remote.OnButtonWasPushed(0)
	_ = remote.OnButtonWasPushed(2)
	_ = remote.UndoButtonWasPushed()
}
