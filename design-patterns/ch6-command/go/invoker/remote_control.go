package invoker

import (
	"command/command"
	"errors"
	"fmt"
)

const slotCount = 7

type RemoteControl struct {
	onCommands  []command.Command
	offCommands []command.Command
	undoCommand command.Command
}

func NewRemoteControl() *RemoteControl {
	var onCommands []command.Command
	var offCommands []command.Command
	for i := 0; i < slotCount; i++ {
		onCommands = append(onCommands, command.NewEmptyCommand())
		offCommands = append(offCommands, command.NewEmptyCommand())
	}
	undoCommand := command.NewEmptyCommand()

	return &RemoteControl{
		onCommands:  onCommands,
		offCommands: offCommands,
		undoCommand: undoCommand,
	}
}

func (r *RemoteControl) SetOnCommand(slot int, command command.Command) (err error) {
	err = validateSlot(slot)
	if err != nil {
		return
	}
	fmt.Println(fmt.Sprintf("set on command button %v", slot))
	r.onCommands[slot] = command
	return
}

func (r *RemoteControl) SetOffCommand(slot int, command command.Command) (err error) {
	err = validateSlot(slot)
	if err != nil {
		return
	}
	fmt.Println(fmt.Sprintf("set off command button %v", slot))
	r.offCommands[slot] = command
	return
}

func (r *RemoteControl) OnButtonWasPushed(slot int) (err error) {
	err = validateSlot(slot)
	if err != nil {
		return
	}
	fmt.Println(fmt.Sprintf("push on command button %v", slot))
	r.onCommands[slot].Execute()
	r.undoCommand = r.onCommands[slot]
	return
}

func (r *RemoteControl) OffButtonWasPushed(slot int) (err error) {
	err = validateSlot(slot)
	if err != nil {
		return
	}
	fmt.Println(fmt.Sprintf("push off command button %v", slot))
	r.offCommands[slot].Execute()
	r.undoCommand = r.offCommands[slot]
	return
}

func (r *RemoteControl) UndoButtonWasPushed() (err error) {
	r.undoCommand.Undo()
	return
}

func validateSlot(slot int) (err error) {
	if slot < 0 || slot >= slotCount {
		err = errors.New("invalid slot")
		fmt.Println(err)
	}
	return
}
