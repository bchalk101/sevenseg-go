package sevenseg

import (
	"github.com/stianeikeland/go-rpio"
	"fmt"
)

const SIZE = 2

type Display struct {
	segment          sevenSegDisplay
	segmentActivePin [SIZE]RaspberryPiPin
}

func NewDisplay(pinA int, pinB int, pinC int, pinD int, pinE int, pinF int, pinG int, pinH int, pinD4 int, pinD3 int) *Display {
	display := new(Display)

	display.segment = NewSevenSeg(pinA, pinB, pinC, pinD, pinE, pinF, pinG, pinH)

	display.segmentActivePin[0] = NewRaspberryPiPin(pinD4)
	display.segmentActivePin[1] = NewRaspberryPiPin(pinD3)

	return display
}

func (d *Display) Print(number string) {
	numberChannel := make(chan string, 1)
	go d.display(numberChannel)
	numberChannel <- number
}

func (d *Display) display(numberChannel chan string) {
	var number string
	for {
		select {
		case number = <-numberChannel:
		default:
		}
		fmt.Printf("Displaying %v\n", number)

		for i := range d.segmentActivePin {
			d.segmentActivePin[i].WriteState(rpio.High)
			d.segment.Display(CLEAR)
		}
		for i, num := range number {
			d.chooseSegment(i)
			d.segment.Display(string(num))
		}
	}

}

func (d *Display) chooseSegment(segment int) {
	for i := range d.segmentActivePin {
		if i == segment {
			d.segmentActivePin[i].WriteState(rpio.High)
		} else {
			d.segmentActivePin[i].WriteState(rpio.Low)
		}
	}
}
