package sevenseg

import  "github.com/stianeikeland/go-rpio"

type sevenSegDisplay interface {
	Display(item string) error
}

type sevenSeg struct {
	toDisplay string
	pinH RaspberryPiPin
	pinD4 RaspberryPiPin
}

const (
	DOT = "."
	ONE = "1"
)


func (s *sevenSeg) Display(item string) error {
	s.toDisplay = item
	s.pinH.WriteState(rpio.High)
	s.pinD4.WriteState(rpio.High)
	return nil
}

func (s *sevenSeg) pinHState() rpio.State {
	return s.pinH.ReadState()
}