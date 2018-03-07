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
	pinD4     RaspberryPiPin
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
)

func NewSevenSeg(pinA int, pinB int, pinC int, pinD int, pinE int, pinF int, pinG int, pinH int, pinD4 int) *sevenSeg {

	display := new(sevenSeg)
	display.pinA = NewRaspberryPiPin(pinA)
	display.pinB = NewRaspberryPiPin(pinB)
	display.pinC = NewRaspberryPiPin(pinC)
	display.pinD = NewRaspberryPiPin(pinD)
	display.pinE = NewRaspberryPiPin(pinE)
	display.pinF = NewRaspberryPiPin(pinF)
	display.pinG = NewRaspberryPiPin(pinG)
	display.pinH = NewRaspberryPiPin(pinH)
	display.pinD4 = NewRaspberryPiPin(pinD4)


	display.pinA.SetMode(rpio.Output)
	display.pinB.SetMode(rpio.Output)
	display.pinC.SetMode(rpio.Output)
	display.pinD.SetMode(rpio.Output)
	display.pinE.SetMode(rpio.Output)
	display.pinF.SetMode(rpio.Output)
	display.pinG.SetMode(rpio.Output)
	display.pinH.SetMode(rpio.Output)
	display.pinD4.SetMode(rpio.Output)


	return display
}

func (s *sevenSeg) Display(item string) error {
	s.toDisplay = item
	s.pinD4.WriteState(rpio.High)

	switch s.toDisplay {
	case DOT:
		s.clear()
		s.pinH.WriteState(rpio.Low)
	case ZERO:
		s.clear()
		s.writeAZero()
	case ONE:
		s.clear()
		s.writeAOne()
	case TWO:
		s.clear()
		s.writeATwo()
	case THREE:
		s.clear()
		s.writeThree()
	case FOUR:
		s.clear()
		s.writeAFour()
	case FIVE:
		s.clear()
		s.writeAFive()
	case SIX:
		s.clear()
		s.writeASix()
	case SEVEN:
		s.clear()
		s.writeASeven()
	case EIGHT:
		s.clear()
		s.writeAEight()
	case NINE:
		s.clear()
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
}
func (s *sevenSeg) writeAFive() {
	s.pinA.WriteState(rpio.Low)
	s.pinC.WriteState(rpio.Low)
	s.pinD.WriteState(rpio.Low)
	s.pinF.WriteState(rpio.Low)
	s.pinG.WriteState(rpio.Low)
}
func (s *sevenSeg) writeAFour() {
	s.pinB.WriteState(rpio.Low)
	s.pinC.WriteState(rpio.Low)
	s.pinF.WriteState(rpio.Low)
	s.pinG.WriteState(rpio.Low)
}

func (s *sevenSeg) writeAOne() {
	s.pinB.WriteState(rpio.Low)
	s.pinC.WriteState(rpio.Low)
}
func (s *sevenSeg) writeATwo() {
	s.pinA.WriteState(rpio.Low)
	s.pinB.WriteState(rpio.Low)
	s.pinD.WriteState(rpio.Low)
	s.pinE.WriteState(rpio.Low)
	s.pinG.WriteState(rpio.Low)
}

func (s *sevenSeg) writeThree() {
	s.pinA.WriteState(rpio.Low)
	s.pinB.WriteState(rpio.Low)
	s.pinC.WriteState(rpio.Low)
	s.pinD.WriteState(rpio.Low)
	s.pinG.WriteState(rpio.Low)
}

func (s *sevenSeg) writeASix() {
	s.pinA.WriteState(rpio.Low)
	s.pinC.WriteState(rpio.Low)
	s.pinD.WriteState(rpio.Low)
	s.pinE.WriteState(rpio.Low)
	s.pinF.WriteState(rpio.Low)
	s.pinG.WriteState(rpio.Low)
}

func (s *sevenSeg) writeASeven() {
	s.pinA.WriteState(rpio.Low)
	s.pinB.WriteState(rpio.Low)
	s.pinC.WriteState(rpio.Low)
}

func (s *sevenSeg) writeANine() {
	s.pinA.WriteState(rpio.Low)
	s.pinB.WriteState(rpio.Low)
	s.pinC.WriteState(rpio.Low)
	s.pinF.WriteState(rpio.Low)
	s.pinG.WriteState(rpio.Low)
}
func (s *sevenSeg) writeAEight() {
	s.pinA.WriteState(rpio.Low)
	s.pinB.WriteState(rpio.Low)
	s.pinC.WriteState(rpio.Low)
	s.pinD.WriteState(rpio.Low)
	s.pinE.WriteState(rpio.Low)
	s.pinF.WriteState(rpio.Low)
	s.pinG.WriteState(rpio.Low)
}
