package main

import (
	"fmt"
	"observer/display"
	"observer/subject"
)

func main() {
	weatherData := subject.NewWeatherData()

	currentDisplay := display.NewCurrentConditionsDisplay(weatherData)
	statisticsDisplay := display.NewStatisticsDisplay(weatherData)
	forecastDisplay := display.NewForecastDisplay(weatherData)
	heatIndexDisplay := display.NewHeatIndexDisplay(weatherData)
	fmt.Println(currentDisplay, statisticsDisplay, forecastDisplay, heatIndexDisplay)

	fmt.Println("\nMeasurements #1")
	weatherData.SetMeasurements(80, 65, 30.4)

	fmt.Println("\nMeasurements #2")
	weatherData.RemoveObserver(forecastDisplay)
	weatherData.SetMeasurements(82, 70, 29.2)

	fmt.Println("\nMeasurements #3")
	weatherData.SetMeasurements(78, 90, 29.2)
}
