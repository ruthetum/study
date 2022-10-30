package component

const (
	TallSize = iota
	GrandeSize
	VentiSize
)

var BeverageSizeMap = map[int]string{
	TallSize:   "TALL",
	GrandeSize: "GRANDE",
	VentiSize:  "VENTI",
}

type Beverage interface {
	GetDescription() string
	Cost() float64
	GetSize() string
	SetSize(size int)
}
