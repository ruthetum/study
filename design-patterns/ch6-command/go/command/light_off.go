package command

import (
	"command/receiver/interface"
)

type LightOffCommand struct {
	light receiver.Light
}

func NewLightOffCommand(light receiver.Light) *LightOffCommand {
	return &LightOffCommand{
		light: light,
	}
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

func (c *LightOffCommand) Undo() {
	c.light.On()
}
