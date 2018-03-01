package sevenseg

import "github.com/stianeikeland/go-rpio"

type RaspberryPiPin interface {
	ReadState() rpio.State
	WriteState(state rpio.State)
}

type raspberryPiPinImpl struct {
	rpioPin rpio.Pin
}

func NewRaspberryPiPin(pin int) RaspberryPiPin {
	return &raspberryPiPinImpl{
		rpioPin: rpio.Pin(pin),
	}
}

func (r *raspberryPiPinImpl) ReadState() rpio.State {
	return r.rpioPin.Read()
}

func (r *raspberryPiPinImpl) WriteState(state rpio.State) {
	r.rpioPin.Write(state)
}
