package main

import (
	"fmt"
	"observer/observer"
	"observer/subject"
)

func main() {
	weatherData := subject.NewWeatherData()

	currentDisplay := observer.NewCurrentConditionsDisplay()
	statisticsDisplay := observer.NewStatisticsDisplay()
	forecastDisplay := observer.NewForecastDisplay()
	heatIndexDisplay := observer.NewHeatIndexDisplay()

	weatherData.RegisterObserver(currentDisplay)
	weatherData.RegisterObserver(statisticsDisplay)
	weatherData.RegisterObserver(forecastDisplay)
	weatherData.RegisterObserver(heatIndexDisplay)

	fmt.Println("\nMeasurements #1")
	weatherData.SetMeasurements(80, 65, 30.4)

	fmt.Println("\nMeasurements #2")
	weatherData.RemoveObserver(forecastDisplay)
	weatherData.SetMeasurements(82, 70, 29.2)

	fmt.Println("\nMeasurements #3")
	weatherData.SetMeasurements(78, 90, 29.2)
}
