package gpio

import "github.com/stianeikeland/go-rpio/v4"

type Action string

// define gpio action
const (
	ActionHigh   = "high"
	ActionLow    = "low"
	ActionToggle = "toggle"
)

func Init() error {
	return rpio.Open()
}

func Release() error {
	return rpio.Close()
}

//https://datasheets.raspberrypi.com/rpi3/raspberry-pi-3-b-plus-reduced-schematics.pdf
func SetAction(p int, action Action ) {
	pin := rpio.Pin(uint8(p))
	pin.Output()
	switch action {
	case ActionHigh:
		pin.High()
	case ActionLow:
		pin.Low()
	case ActionToggle:
		pin.Toggle()
	}
}
