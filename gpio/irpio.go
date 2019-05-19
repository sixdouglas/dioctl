package gpio

import (
	"github.com/stianeikeland/go-rpio"
)

type IRpio interface {
	PinMode(pin rpio.Pin, mode rpio.Mode)
	WritePin(pin rpio.Pin, state rpio.State)
	ReadPin(pin rpio.Pin) rpio.State
	TogglePin(pin rpio.Pin)
	DetectEdge(pin rpio.Pin, edge rpio.Edge)
	EdgeDetected(pin rpio.Pin) bool
	PullMode(pin rpio.Pin, pull rpio.Pull)
	SetFreq(pin rpio.Pin, freq int)
	SetDutyCycle(pin rpio.Pin, dutyLen, cycleLen uint32)
	StopPwm()
	StartPwm()
	EnableIRQs(irqs uint64)
	DisableIRQs(irqs uint64)
	Open() (err error)
	Close() (err error)
}

type RpioObj struct {}

func (r RpioObj) PinMode(pin rpio.Pin, mode rpio.Mode) {
	rpio.PinMode(pin, mode)
}

func (r RpioObj) WritePin(pin rpio.Pin, state rpio.State) {
	rpio.WritePin(pin, state)
}

func (r RpioObj) ReadPin(pin rpio.Pin) rpio.State {
	return rpio.ReadPin(pin)
}

func (r RpioObj) TogglePin(pin rpio.Pin) {
	rpio.TogglePin(pin)
}

func (r RpioObj) DetectEdge(pin rpio.Pin, edge rpio.Edge) {
	rpio.DetectEdge(pin, edge)
}

func (r RpioObj) EdgeDetected(pin rpio.Pin) bool {
	return rpio.EdgeDetected(pin)
}

func (r RpioObj) PullMode(pin rpio.Pin, pull rpio.Pull) {
	rpio.PullMode(pin, pull)
}

func (r RpioObj) SetFreq(pin rpio.Pin, freq int) {
	rpio.SetFreq(pin, freq)
}

func (r RpioObj) SetDutyCycle(pin rpio.Pin, dutyLen, cycleLen uint32) {
	rpio.SetDutyCycle(pin, dutyLen, cycleLen)
}

func (r RpioObj) StopPwm() {
	rpio.StopPwm()
}

func (r RpioObj) StartPwm() {
	rpio.StartPwm()
}

func (r RpioObj) EnableIRQs(irqs uint64) {
	rpio.EnableIRQs(irqs)
}

func (r RpioObj) DisableIRQs(irqs uint64) {
	rpio.DisableIRQs(irqs)
}

func (r RpioObj) Close() (err error) {
	return rpio.Close()
}

func (r RpioObj) Open() (err error) {
	return rpio.Open()
}
