package observable

import "observer/observer"

type WeatherData struct {
	Observers   []observer.Observer
	Temperature float64
	Humidity    float64
	Pressure    float64
}

func NewWeatherData() *WeatherData {
	return &WeatherData{
		Observers: []observer.Observer{},
	}
}

func (d *WeatherData) RegisterObserver(observer observer.Observer) {
	d.Observers = append(d.Observers, observer)
}

func (d *WeatherData) RemoveObserver(observer observer.Observer) {
	for k, v := range d.Observers {
		if v == observer {
			d.Observers = remove(d.Observers, k)
		}
	}
}

func (d *WeatherData) NotifyAll() {
	for _, o := range d.Observers {
		o.Update(d.Temperature, d.Humidity, d.Pressure)
	}
}

func (d *WeatherData) MeasurementsChanged() {
	d.NotifyAll()
}

func (d *WeatherData) SetMeasurements(temp, humidity, pressure float64) {
	d.Temperature = temp
	d.Humidity = humidity
	d.Pressure = pressure
	d.MeasurementsChanged()
}

func (d *WeatherData) GetTemperature() float64 {
	return d.Temperature
}

func (d *WeatherData) GetHumidity() float64 {
	return d.Humidity
}

func (d *WeatherData) GetPressure() float64 {
	return d.Pressure
}

func remove(slice []observer.Observer, index int) []observer.Observer {
	return append(slice[:index], slice[index+1:]...)
}
