package controller

import (
	"bridge/tv"
	"fmt"
)

type SonyRemoteControl struct {
	tv tv.TV
}

func NewSonyRemoteControl(tv tv.TV) *SonyRemoteControl {
	return &SonyRemoteControl{tv: tv}
}

func (c *SonyRemoteControl) On() {
	fmt.Println("Sony RemoteControl on")
	c.tv.On()
}

func (c *SonyRemoteControl) Off() {
	fmt.Println("Sony RemoteControl off")
	c.tv.Off()
}

func (c *SonyRemoteControl) SetChannel() {
	c.tv.TuneChannel()
}
