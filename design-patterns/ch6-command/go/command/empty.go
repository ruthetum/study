package command

import "fmt"

type EmptyCommand struct {
}

func NewEmptyCommand() *EmptyCommand {
	return &EmptyCommand{}
}

func (c *EmptyCommand) Execute() {
	fmt.Println("nothing happens")
}

func (c *EmptyCommand) Undo() {
	fmt.Println("nothing happens")
}
