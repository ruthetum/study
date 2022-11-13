package receiver

type Stereo interface {
	On()
	Off()
	SetCD()
	SetDVD()
	SetRadio()
	SetVolume(value int)
	GetVolume() int
}
