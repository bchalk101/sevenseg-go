package sevenseg

import (
	"github.com/stianeikeland/go-rpio"
	"fmt"
)

const SIZE = 2

type Display struct {
	segment          sevenSegDisplay
	segmentActivePin [SIZE]RaspberryPiPin
	displayStarted   bool
	numberChannel    chan string
}

func NewDisplay(pinA int, pinB int, pinC int, pinD int, pinE int, pinF int, pinG int, pinH int, pinD4 int, pinD3 int) *Display {
	display := new(Display)

	display.segment = NewSevenSeg(pinA, pinB, pinC, pinD, pinE, pinF, pinG, pinH)

	display.segmentActivePin[0] = NewRaspberryPiPin(pinD4)
	display.segmentActivePin[1] = NewRaspberryPiPin(pinD3)

	display.setupDisplayChannel()

	return display
}
func (display *Display) setupDisplayChannel() {
	display.numberChannel = make(chan string, 1)
	display.numberChannel <- "00"
	go display.display()
}

func (d *Display) Print(number string) {
	d.numberChannel <- number
}

func (d *Display) display() {
	var number string
	for {

		select {
		case number = <-d.numberChannel:
			fmt.Printf("Changing to %v\n", number)
			for i := range d.segmentActivePin {
				d.segmentActivePin[i].WriteState(rpio.High)
				d.segment.Display(CLEAR)
			}
		default:
		}
		for i, num := range number {
			d.deselectAllSegments()
			d.segment.Display(CLEAR)
			d.chooseSegment(i)
			d.segment.Display(string(num))
		}
	}

}

func (d *Display) deselectAllSegments() {
	for i := range d.segmentActivePin {
		d.segmentActivePin[i].WriteState(rpio.Low)
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
