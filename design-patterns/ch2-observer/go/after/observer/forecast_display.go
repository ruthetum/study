package observer

import (
	"fmt"
)

type ForecastDisplay struct {
	CurrentPressure float64
	LastPressure    float64
}

func NewForecastDisplay() *ForecastDisplay {
	display := &ForecastDisplay{}
	return display
}

func (d *ForecastDisplay) Update(temp, humidity, pressure float64) {
	d.LastPressure = d.CurrentPressure
	d.CurrentPressure = pressure
	d.Display()
}

func (d *ForecastDisplay) Display() {
	fmt.Println("[Forecast]")
	if d.CurrentPressure > d.LastPressure {
		fmt.Println("Improving weather on the way!")
	} else if d.CurrentPressure == d.LastPressure {
		fmt.Println("More of the same")
	} else {
		fmt.Println("Watch out for cooler, rainy weather")
	}
}
