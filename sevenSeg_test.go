package sevenseg

import (
	"testing"
	"github.com/stianeikeland/go-rpio"
)

var displayTests = []struct {
	display string
	pinAState rpio.State
	pinBState rpio.State
	pinHState rpio.State
	pinD4State rpio.State
} {
	{DOT, rpio.Low,rpio.Low, rpio.High, rpio.High},
	{ONE, rpio.High, rpio.High, rpio.Low, rpio.High},
}

func TestDisplay(t *testing.T) {
	sevenSeg := new(sevenSeg)
	sevenSeg.pinA = new(mockRaspberryPiPinImpl)
	sevenSeg.pinB = new(mockRaspberryPiPinImpl)
	sevenSeg.pinH = new(mockRaspberryPiPinImpl)
	sevenSeg.pinD4 = new(mockRaspberryPiPinImpl)

	for _, testState := range displayTests {
		sevenSeg.Display(testState.display)


		if sevenSeg.toDisplay != testState.display {
			t.Errorf("Not displaying correct item: Expected %s but got %s", testState.display ,sevenSeg.toDisplay)
		}

		if sevenSeg.pinA.ReadState() != testState.pinAState  {
			t.Errorf("Display %v: Expected Pin A to be %v but was %v", testState.display, testState.pinAState, sevenSeg.pinA.ReadState())
		}

		if sevenSeg.pinB.ReadState() != testState.pinBState  {
			t.Errorf("Display %v: Expected Pin B to be %v but was %v", testState.display, testState.pinBState, sevenSeg.pinB.ReadState())
		}

		if sevenSeg.pinH.ReadState() != testState.pinHState  {
			t.Errorf("Display %v: Expected Pin H to be %v but was %v", testState.display, testState.pinHState, sevenSeg.pinH.ReadState())
		}

		if sevenSeg.pinD4.ReadState() != testState.pinD4State {
			t.Errorf("Display %v: Expected Pin D4 to be %v but was %v", testState.display, testState.pinD4State, sevenSeg.pinD4.ReadState())
		}
	}
}
