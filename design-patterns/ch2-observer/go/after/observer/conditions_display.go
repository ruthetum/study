package observer

import (
	"fmt"
)

type CurrentConditionsDisplay struct {
	Temperature float64
	Humidity    float64
}

func NewCurrentConditionsDisplay() *CurrentConditionsDisplay {
	display := &CurrentConditionsDisplay{}
	return display
}

func (d *CurrentConditionsDisplay) Update(temp, humidity, pressure float64) {
	d.Temperature = temp
	d.Humidity = humidity
	d.Display()
}

func (d *CurrentConditionsDisplay) Display() {
	fmt.Println("[Current conditions]")
	fmt.Println(fmt.Sprintf("Temperature: %v", d.Temperature))
	fmt.Println(fmt.Sprintf("Humidity: %v", d.Humidity))
}
