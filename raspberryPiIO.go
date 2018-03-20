package sevenseg

import "github.com/stianeikeland/go-rpio"

type RaspberryPiPin interface {
	ReadState() rpio.State
	WriteState(state rpio.State)
	SetMode(mode rpio.Mode)

}

type raspberryPiPinImpl struct {
	rpioPin rpio.Pin
}

func NewRaspberryPiPin(pin int) RaspberryPiPin {
	raspberryPiPin := new(raspberryPiPinImpl)
	raspberryPiPin.rpioPin = rpio.Pin(pin)
	raspberryPiPin.SetMode(rpio.Output)
	return raspberryPiPin
}

func (r *raspberryPiPinImpl) ReadState() rpio.State {
	return r.rpioPin.Read()
}

func (r *raspberryPiPinImpl) WriteState(state rpio.State) {
	r.rpioPin.Write(state)
}

func (r *raspberryPiPinImpl) SetMode(mode rpio.Mode) {
	r.rpioPin.Mode(mode)
}