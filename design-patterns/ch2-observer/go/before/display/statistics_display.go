package display

import (
	"fmt"
	"observer/subject"
)

type StatisticsDisplay struct {
	MaxTemperature float64
	MinTemperature float64
	TemperatureSum float64
	NumReadings    int
	Subject        subject.Subject
}

func NewStatisticsDisplay(subject subject.Subject) *StatisticsDisplay {
	display := &StatisticsDisplay{
		Subject: subject,
	}
	display.MinTemperature = 200
	subject.RegisterObserver(display)
	return display
}

func (d *StatisticsDisplay) Update(temp, humidity, pressure float64) {
	d.TemperatureSum += temp
	d.NumReadings++

	if temp > d.MaxTemperature {
		d.MaxTemperature = temp
	}
	if temp < d.MinTemperature {
		d.MinTemperature = temp
	}

	d.Display()
}

func (d *StatisticsDisplay) Display() {
	fmt.Println("[Statistics]")
	fmt.Println(fmt.Sprintf("Avg: %v", d.TemperatureSum/float64(d.NumReadings)))
	fmt.Println(fmt.Sprintf("Max: %v", d.MaxTemperature))
	fmt.Println(fmt.Sprintf("Min: %v", d.MinTemperature))
}
