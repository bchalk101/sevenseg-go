package sevenseg

const SIZE = 2

type Display struct {
	segment [SIZE]sevenSegDisplay
}

func NewDisplay(pinA int, pinB int, pinC int, pinD int, pinE int, pinF int, pinG int, pinH int, pinD4 int, pinD3 int) *Display {
	display := new(Display)

	display.segment[0] = NewSevenSeg(pinA, pinB, pinC, pinD, pinE, pinF, pinG, pinH, pinD4)
	display.segment[1] = NewSevenSeg(pinA, pinB, pinC, pinD, pinE, pinF, pinG, pinH, pinD3)

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
	for i, num := range number {
		d.segment[i].Display(string(num))
	}
}
