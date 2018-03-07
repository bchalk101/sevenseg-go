package sevenseg

import (
	"testing"
	"github.com/stianeikeland/go-rpio"
)

var displayTests = []struct {
	display    string
	pinAState  rpio.State
	pinBState  rpio.State
	pinCState  rpio.State
	pinDState  rpio.State
	pinEState  rpio.State
	pinFState  rpio.State
	pinGState  rpio.State
	pinHState  rpio.State
	pinD4State rpio.State
}{
	{DOT, rpio.High, rpio.High, rpio.High, rpio.High, rpio.High, rpio.High, rpio.High, rpio.Low, rpio.High},
	{ZERO, rpio.Low, rpio.Low, rpio.Low, rpio.Low, rpio.Low, rpio.Low, rpio.High, rpio.High, rpio.High},
	{ONE, rpio.High, rpio.Low, rpio.Low, rpio.High, rpio.High, rpio.High, rpio.High, rpio.High, rpio.High},
	{TWO, rpio.Low, rpio.Low, rpio.High, rpio.Low, rpio.Low, rpio.High, rpio.Low, rpio.High, rpio.High},
	{THREE, rpio.Low, rpio.Low, rpio.Low, rpio.Low, rpio.High, rpio.High, rpio.Low, rpio.High, rpio.High},
	{FOUR, rpio.High, rpio.Low, rpio.Low, rpio.High, rpio.High, rpio.Low, rpio.Low, rpio.High, rpio.High},
	{FIVE, rpio.Low, rpio.High, rpio.Low, rpio.Low, rpio.High, rpio.Low, rpio.Low, rpio.High, rpio.High},
	{SIX, rpio.Low, rpio.High, rpio.Low, rpio.Low, rpio.Low, rpio.Low, rpio.Low, rpio.High, rpio.High},
	{SEVEN, rpio.Low, rpio.Low, rpio.Low, rpio.High, rpio.High, rpio.High, rpio.High, rpio.High, rpio.High},
	{EIGHT, rpio.Low, rpio.Low, rpio.Low, rpio.Low, rpio.Low, rpio.Low, rpio.Low, rpio.High, rpio.High},
	{NINE, rpio.Low, rpio.Low, rpio.Low, rpio.High, rpio.High, rpio.Low, rpio.Low, rpio.High, rpio.High},
}

func TestDisplay(t *testing.T) {
	sevenSeg := new(sevenSeg)
	sevenSeg.pinA = new(mockRaspberryPiPinImpl)
	sevenSeg.pinB = new(mockRaspberryPiPinImpl)
	sevenSeg.pinC = new(mockRaspberryPiPinImpl)
	sevenSeg.pinD = new(mockRaspberryPiPinImpl)
	sevenSeg.pinE = new(mockRaspberryPiPinImpl)
	sevenSeg.pinF = new(mockRaspberryPiPinImpl)
	sevenSeg.pinG = new(mockRaspberryPiPinImpl)
	sevenSeg.pinH = new(mockRaspberryPiPinImpl)
	sevenSeg.pinD4 = new(mockRaspberryPiPinImpl)

	for _, testState := range displayTests {
		sevenSeg.Display(testState.display)

		if sevenSeg.toDisplay != testState.display {
			t.Errorf("Not displaying correct item: Expected %s but got %s", testState.display, sevenSeg.toDisplay)
		}

		if sevenSeg.pinA.ReadState() != testState.pinAState {
			t.Errorf("Display %v: Expected Pin A to be %v but was %v", testState.display, testState.pinAState, sevenSeg.pinA.ReadState())
		}

		if sevenSeg.pinB.ReadState() != testState.pinBState {
			t.Errorf("Display %v: Expected Pin B to be %v but was %v", testState.display, testState.pinBState, sevenSeg.pinB.ReadState())
		}
		if sevenSeg.pinC.ReadState() != testState.pinCState {
			t.Errorf("Display %v: Expected Pin C to be %v but was %v", testState.display, testState.pinCState, sevenSeg.pinC.ReadState())
		}
		if sevenSeg.pinD.ReadState() != testState.pinDState {
			t.Errorf("Display %v: Expected Pin D to be %v but was %v", testState.display, testState.pinDState, sevenSeg.pinD.ReadState())
		}
		if sevenSeg.pinE.ReadState() != testState.pinEState {
			t.Errorf("Display %v: Expected Pin E to be %v but was %v", testState.display, testState.pinEState, sevenSeg.pinE.ReadState())
		}
		if sevenSeg.pinF.ReadState() != testState.pinFState {
			t.Errorf("Display %v: Expected Pin F to be %v but was %v", testState.display, testState.pinFState, sevenSeg.pinF.ReadState())
		}
		if sevenSeg.pinG.ReadState() != testState.pinGState {
			t.Errorf("Display %v: Expected Pin G to be %v but was %v", testState.display, testState.pinGState, sevenSeg.pinG.ReadState())
		}

		if sevenSeg.pinH.ReadState() != testState.pinHState {
			t.Errorf("Display %v: Expected Pin H to be %v but was %v", testState.display, testState.pinHState, sevenSeg.pinH.ReadState())
		}

		if sevenSeg.pinD4.ReadState() != testState.pinD4State {
			t.Errorf("Display %v: Expected Pin D4 to be %v but was %v", testState.display, testState.pinD4State, sevenSeg.pinD4.ReadState())
		}
	}
}
