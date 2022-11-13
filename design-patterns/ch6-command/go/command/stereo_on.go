package command

import (
	receiver "command/receiver/interface"
)

type StereoOnWithCDCommand struct {
	stereo     receiver.Stereo
	prevVolume int
}

func NewStereoOnWithCDCommand(stereo receiver.Stereo) *StereoOnWithCDCommand {
	return &StereoOnWithCDCommand{
		stereo: stereo,
	}
}

func (c *StereoOnWithCDCommand) Execute() {
	c.stereo.On()
	c.stereo.SetCD()
	c.prevVolume = c.stereo.GetVolume()
	c.stereo.SetVolume(11)
}

func (c *StereoOnWithCDCommand) Undo() {
	c.stereo.SetVolume(c.prevVolume)
	c.stereo.Off()
}
