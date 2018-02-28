package sevenseg

import "github.com/stianeikeland/go-rpio"

type mockRaspberryPiPinImpl struct {
	stateOfPin rpio.State
	readStateCalled bool
	writeStateCalled bool
}


func NewMockRaspberryPiPin(pin int) mockRaspberryPiPinImpl {
	return mockRaspberryPiPinImpl{

	}
}

func (r *mockRaspberryPiPinImpl) ReadState() rpio.State {
	return r.stateOfPin
}

func (r *mockRaspberryPiPinImpl) WriteState(state rpio.State) {
	r.stateOfPin = state
	r.writeStateCalled = true
}