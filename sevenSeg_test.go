package sevenseg

import (
	"testing"
	"github.com/stianeikeland/go-rpio"
)

var displayTests = []struct {
	display string
	pinHState rpio.State
	pinD4State rpio.State
} {
	{DOT, rpio.High, rpio.High},
}

func TestDisplay(t *testing.T) {
	sevenSeg := new(sevenSeg)
	sevenSeg.pinH = new(mockRaspberryPiPinImpl)
	sevenSeg.pinD4 = new(mockRaspberryPiPinImpl)

	for _, testState := range displayTests {
		sevenSeg.Display(testState.display)


		if sevenSeg.toDisplay != testState.display {
			t.Errorf("Not displaying correct item: Expected %s but got %s", testState.display ,sevenSeg.toDisplay)
		}

		if sevenSeg.pinH.ReadState() != testState.pinHState  {
			t.Errorf("Expected Pin H to be %vbut was %s", testState.pinHState, sevenSeg.pinH.ReadState())
		}

		if sevenSeg.pinD4.ReadState() != testState.pinD4State {
			t.Errorf("Expected Pin D4 to be %v but was %v", testState.pinD4State, sevenSeg.pinD4.ReadState())
		}
	}
}
