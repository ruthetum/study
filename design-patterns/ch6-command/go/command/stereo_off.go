package command

import receiver "command/receiver/interface"

type StereoOffWithCDCommand struct {
	stereo     receiver.Stereo
	prevVolume int
}

func NewStereoOffWithCDCommand(stereo receiver.Stereo) *StereoOffWithCDCommand {
	return &StereoOffWithCDCommand{
		stereo: stereo,
	}
}

func (c *StereoOffWithCDCommand) Execute() {
	c.prevVolume = c.stereo.GetVolume()
	c.stereo.Off()
}

func (c *StereoOffWithCDCommand) Undo() {
	c.stereo.On()
	c.stereo.SetVolume(c.prevVolume)
}
