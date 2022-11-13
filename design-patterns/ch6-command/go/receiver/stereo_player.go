package receiver

import "fmt"

type StereoPlayer struct {
	volume int
}

func NewStereoPlayer() *StereoPlayer {
	return &StereoPlayer{
		volume: 0,
	}
}

func (s *StereoPlayer) On() {
	fmt.Println("Turning on StereoPlayer")
}

func (s *StereoPlayer) Off() {
	fmt.Println("Turning off StereoPlayer")
}

func (s *StereoPlayer) SetCD() {
	fmt.Println("Setting CD on StereoPlayer")
}

func (s *StereoPlayer) SetDVD() {
	fmt.Println("Setting DVD on StereoPlayer")
}

func (s *StereoPlayer) SetRadio() {
	fmt.Println("Setting Radio on StereoPlayer")
}

func (s *StereoPlayer) SetVolume(value int) {
	fmt.Println(fmt.Sprintf("Setting volume %v on StereoPlayer", value))
	s.volume = value
}

func (s *StereoPlayer) GetVolume() (volume int) {
	return s.volume
}
