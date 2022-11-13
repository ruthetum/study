package command

import (
	"command/receiver/interface"
)

type LightOnCommand struct {
	light receiver.Light
}

func NewLightOnCommand(light receiver.Light) *LightOnCommand {
	return &LightOnCommand{
		light: light,
	}
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

func (c *LightOnCommand) Undo() {
	c.light.Off()
}
