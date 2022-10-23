package display

import (
	"fmt"
	"observer/subject"
)

type ForecastDisplay struct {
	CurrentPressure float64
	LastPressure    float64
	Subject         subject.Subject
}

func NewForecastDisplay(subject subject.Subject) *ForecastDisplay {
	display := &ForecastDisplay{
		Subject: subject,
	}
	subject.RegisterObserver(display)
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
