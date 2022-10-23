package observer

type Observer interface {
	Update(temp float64, humidity float64, pressure float64)
}
