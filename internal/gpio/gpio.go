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

func SetAction(pin int, action Action ) {
	upin := uint8(pin)
	rpio.Pin(pin).Output()
	switch action {
	case ActionHigh:
		rpio.Pin(upin).High()
	case ActionLow:
		rpio.Pin(upin).Low()
	case ActionToggle:
		rpio.Pin(upin).Toggle()
	}
}
