package gpio

import "github.com/stianeikeland/go-rpio"

type IDio interface {
	PulseIn(rpio IRpio, pin rpio.Pin, state rpio.State, timeout int64) (int64, error)
	ReadCode(rpio IRpio, pin rpio.Pin) (uint64, error)
	SendCommand(rpio IRpio, pin rpio.Pin, senderId uint64, interrupterId uint64, on bool) error
}

type DioObj struct {}

func (r DioObj) PulseIn(rpio IRpio, pin rpio.Pin, state rpio.State, timeout int64) (int64, error) {
	return PulseIn(rpio, pin, state, timeout)
}

func (r DioObj) ReadCode(rpio IRpio, pin rpio.Pin) (uint64, error) {
	return ReadCode(rpio, pin)
}

func (r DioObj) SendCommand(rpio IRpio, pin rpio.Pin, senderId uint64, interrupterId uint64, on bool) {
	SendCommand(rpio, pin, senderId, interrupterId, on)
}
