package gpio

import (
	"github.com/stianeikeland/go-rpio"
)

type RpioObj struct {
	IsOpen bool
}

func (r *RpioObj) PinMode(pin rpio.Pin, mode rpio.Mode) {
	rpio.PinMode(pin, mode)
}

func (r *RpioObj) WritePin(pin rpio.Pin, state rpio.State) {
	rpio.WritePin(pin, state)
}

func (r *RpioObj) ReadPin(pin rpio.Pin) rpio.State {
	return rpio.ReadPin(pin)
}

func (r *RpioObj) TogglePin(pin rpio.Pin) {
	rpio.TogglePin(pin)
}

func (r *RpioObj) DetectEdge(pin rpio.Pin, edge rpio.Edge) {
	rpio.DetectEdge(pin, edge)
}

func (r *RpioObj) EdgeDetected(pin rpio.Pin) bool {
	return rpio.EdgeDetected(pin)
}

func (r *RpioObj) PullMode(pin rpio.Pin, pull rpio.Pull) {
	rpio.PullMode(pin, pull)
}

func (r *RpioObj) SetFreq(pin rpio.Pin, freq int) {
	rpio.SetFreq(pin, freq)
}

func (r *RpioObj) SetDutyCycle(pin rpio.Pin, dutyLen, cycleLen uint32) {
	rpio.SetDutyCycle(pin, dutyLen, cycleLen)
}

func (r *RpioObj) StopPwm() {
	rpio.StopPwm()
}

func (r *RpioObj) StartPwm() {
	rpio.StartPwm()
}

func (r *RpioObj) EnableIRQs(irqs uint64) {
	rpio.EnableIRQs(irqs)
}

func (r *RpioObj) DisableIRQs(irqs uint64) {
	rpio.DisableIRQs(irqs)
}

func (r *RpioObj) Close() (err error) {
	r.IsOpen = false
	return rpio.Close()
}

func (r *RpioObj) Open() (err error) {
	r.IsOpen = true
	return rpio.Open()
}
