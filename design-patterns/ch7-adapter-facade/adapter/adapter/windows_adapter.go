package adapter

import (
	"adapter/adaptee"
	"fmt"
)

type WindowsAdapter struct {
	windows adaptee.Windows
}

func NewWindowsAdapter(windows adaptee.Windows) *WindowsAdapter {
	return &WindowsAdapter{
		windows: windows,
	}
}

func (a *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	a.windows.InsertIntoUSBPort()
}
