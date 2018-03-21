package sevenseg

import "github.com/stianeikeland/go-rpio"

type sevenSegDisplay interface {
	Display(item string) error
}

type sevenSeg struct {
	toDisplay string
	pinA      RaspberryPiPin
	pinB      RaspberryPiPin
	pinC      RaspberryPiPin
	pinD      RaspberryPiPin
	pinE      RaspberryPiPin
	pinF      RaspberryPiPin
	pinG      RaspberryPiPin
	pinH      RaspberryPiPin
}

const (
	DOT   = "."
	ZERO  = "0"
	ONE   = "1"
	TWO   = "2"
	THREE = "3"
	FOUR  = "4"
	FIVE  = "5"
	SIX   = "6"
	SEVEN = "7"
	EIGHT = "8"
	NINE  = "9"
	CLEAR = "000000"
)

func NewSevenSeg(pinA int, pinB int, pinC int, pinD int, pinE int, pinF int, pinG int, pinH int) *sevenSeg {

	display := new(sevenSeg)
	display.pinA = NewRaspberryPiPin(pinA)
	display.pinB = NewRaspberryPiPin(pinB)
	display.pinC = NewRaspberryPiPin(pinC)
	display.pinD = NewRaspberryPiPin(pinD)
	display.pinE = NewRaspberryPiPin(pinE)
	display.pinF = NewRaspberryPiPin(pinF)
	display.pinG = NewRaspberryPiPin(pinG)
	display.pinH = NewRaspberryPiPin(pinH)

	return display
}

func (s *sevenSeg) Display(item string) error {
	s.toDisplay = item

	switch s.toDisplay {
	case CLEAR:
		s.clear()
	case DOT:
		s.clear()
		s.pinH.WriteState(rpio.Low)
	case ZERO:
		s.writeAZero()
	case ONE:
		s.writeAOne()
	case TWO:
		s.writeATwo()
	case THREE:
		s.writeThree()
	case FOUR:
		s.writeAFour()
	case FIVE:
		s.writeAFive()
	case SIX:
		s.writeASix()
	case SEVEN:
		s.writeASeven()
	case EIGHT:
		s.writeAEight()
	case NINE:
		s.writeANine()
	}

	return nil
}

func (s *sevenSeg) clear() {
	s.pinA.WriteState(rpio.High)
	s.pinB.WriteState(rpio.High)
	s.pinC.WriteState(rpio.High)
	s.pinD.WriteState(rpio.High)
	s.pinE.WriteState(rpio.High)
	s.pinF.WriteState(rpio.High)
	s.pinG.WriteState(rpio.High)
	s.pinH.WriteState(rpio.High)
}
func (s *sevenSeg) writeAZero() {
	s.pinA.WriteState(rpio.Low)
	s.pinB.WriteState(rpio.Low)
	s.pinC.WriteState(rpio.Low)
	s.pinD.WriteState(rpio.Low)
	s.pinE.WriteState(rpio.Low)
	s.pinF.WriteState(rpio.Low)
	s.pinG.WriteState(rpio.High)
	s.pinH.WriteState(rpio.High)
}
func (s *sevenSeg) writeAFive() {
	s.pinA.WriteState(rpio.Low)
	s.pinB.WriteState(rpio.High)
	s.pinC.WriteState(rpio.Low)
	s.pinD.WriteState(rpio.Low)
	s.pinE.WriteState(rpio.High)
	s.pinF.WriteState(rpio.Low)
	s.pinG.WriteState(rpio.Low)
	s.pinH.WriteState(rpio.High)
}
func (s *sevenSeg) writeAFour() {
	s.pinA.WriteState(rpio.High)
	s.pinB.WriteState(rpio.Low)
	s.pinC.WriteState(rpio.Low)
	s.pinD.WriteState(rpio.High)
	s.pinE.WriteState(rpio.High)
	s.pinF.WriteState(rpio.Low)
	s.pinG.WriteState(rpio.Low)
	s.pinH.WriteState(rpio.High)
}

func (s *sevenSeg) writeAOne() {
	s.pinA.WriteState(rpio.High)
	s.pinB.WriteState(rpio.Low)
	s.pinC.WriteState(rpio.Low)
	s.pinD.WriteState(rpio.High)
	s.pinE.WriteState(rpio.High)
	s.pinF.WriteState(rpio.High)
	s.pinG.WriteState(rpio.High)
	s.pinH.WriteState(rpio.High)
}
func (s *sevenSeg) writeATwo() {
	s.pinA.WriteState(rpio.Low)
	s.pinB.WriteState(rpio.Low)
	s.pinC.WriteState(rpio.High)
	s.pinD.WriteState(rpio.Low)
	s.pinE.WriteState(rpio.Low)
	s.pinF.WriteState(rpio.High)
	s.pinG.WriteState(rpio.Low)
	s.pinH.WriteState(rpio.High)
}

func (s *sevenSeg) writeThree() {
	s.pinA.WriteState(rpio.Low)
	s.pinB.WriteState(rpio.Low)
	s.pinC.WriteState(rpio.Low)
	s.pinD.WriteState(rpio.Low)
	s.pinE.WriteState(rpio.High)
	s.pinF.WriteState(rpio.High)
	s.pinG.WriteState(rpio.Low)
	s.pinH.WriteState(rpio.High)
}

func (s *sevenSeg) writeASix() {
	s.pinA.WriteState(rpio.Low)
	s.pinB.WriteState(rpio.High)
	s.pinC.WriteState(rpio.Low)
	s.pinD.WriteState(rpio.Low)
	s.pinE.WriteState(rpio.Low)
	s.pinF.WriteState(rpio.Low)
	s.pinG.WriteState(rpio.Low)
	s.pinH.WriteState(rpio.High)

}

func (s *sevenSeg) writeASeven() {
	s.pinA.WriteState(rpio.Low)
	s.pinB.WriteState(rpio.Low)
	s.pinC.WriteState(rpio.Low)
	s.pinD.WriteState(rpio.High)
	s.pinE.WriteState(rpio.High)
	s.pinF.WriteState(rpio.High)
	s.pinG.WriteState(rpio.High)
	s.pinH.WriteState(rpio.High)
}

func (s *sevenSeg) writeANine() {
	s.pinA.WriteState(rpio.Low)
	s.pinB.WriteState(rpio.Low)
	s.pinC.WriteState(rpio.Low)
	s.pinD.WriteState(rpio.High)
	s.pinE.WriteState(rpio.High)
	s.pinF.WriteState(rpio.Low)
	s.pinG.WriteState(rpio.Low)
	s.pinH.WriteState(rpio.High)
}
func (s *sevenSeg) writeAEight() {
	s.pinA.WriteState(rpio.Low)
	s.pinB.WriteState(rpio.Low)
	s.pinC.WriteState(rpio.Low)
	s.pinD.WriteState(rpio.Low)
	s.pinE.WriteState(rpio.Low)
	s.pinF.WriteState(rpio.Low)
	s.pinG.WriteState(rpio.Low)
	s.pinH.WriteState(rpio.High)
}

