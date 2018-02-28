package sevenseg

import "github.com/stianeikeland/go-rpio"

type sevenSegDisplay interface {
	Display(item string) error
}

type sevenSeg struct {
	toDisplay string
	pinA      RaspberryPiPin
	pinB      RaspberryPiPin
	pinH      RaspberryPiPin
	pinD4     RaspberryPiPin
}

const (
	DOT = "."
	ONE = "1"
)

func (s *sevenSeg) Display(item string) error {
	s.toDisplay = item
	switch s.toDisplay {
	case DOT:
		s.pinA.WriteState(rpio.Low)
		s.pinB.WriteState(rpio.Low)
		s.pinH.WriteState(rpio.High)
		s.pinD4.WriteState(rpio.High)
	case ONE:
		s.pinA.WriteState(rpio.High)
		s.pinB.WriteState(rpio.High)
		s.pinH.WriteState(rpio.Low)
		s.pinD4.WriteState(rpio.High)
	}

	return nil
}

func (s *sevenSeg) pinHState() rpio.State {
	return s.pinH.ReadState()
}
