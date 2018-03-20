package sevenseg

import "github.com/stianeikeland/go-rpio"

const SIZE = 2

type Display struct {
	segment          sevenSegDisplay
	segmentActivePin [SIZE]RaspberryPiPin
}

func NewDisplay(pinA int, pinB int, pinC int, pinD int, pinE int, pinF int, pinG int, pinH int) *Display {
	display := new(Display)

<<<<<<< Updated upstream
	display.segment = NewSevenSeg(pinA, pinB, pinC, pinD, pinE, pinF, pinG, pinH)

	display.segmentActivePin[0] = NewRaspberryPiPin(pinD4)
	display.segmentActivePin[1] = NewRaspberryPiPin(pinD3)
=======
	display.segment[0] = NewSevenSeg(pinA, pinB, pinC, pinD, pinE, pinF, pinG, pinH)
	display.segment[1] = NewSevenSeg(pinA, pinB, pinC, pinD, pinE, pinF, pinG, pinH)
>>>>>>> Stashed changes

	return display
}

func (d *Display) Print(number string) {
	go func() {
		for {
			d.display(number)
		}
	}()

}

func (d *Display) display(number string) {
	for i := range d.segmentActivePin {
		d.segmentActivePin[i].WriteState(rpio.High)
		d.segment.Display(CLEAR)
	}
	for i, num := range number {
		d.segmentActivePin[i].WriteState(rpio.High)
		d.segment.Display(string(num))
	}
}
