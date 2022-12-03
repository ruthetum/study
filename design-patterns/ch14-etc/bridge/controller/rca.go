package controller

import (
	"bridge/tv"
	"fmt"
)

type RCARemoteControl struct {
	tv tv.TV
}

func NewRCARemoteControl(tv tv.TV) *RCARemoteControl {
	return &RCARemoteControl{tv: tv}
}

func (c *RCARemoteControl) On() {
	fmt.Println("RCA RemoteControl on")
	c.tv.On()
}

func (c *RCARemoteControl) Off() {
	fmt.Println("RCA RemoteControl off")
	c.tv.Off()
}

func (c *RCARemoteControl) SetChannel() {
	c.tv.TuneChannel()
}
