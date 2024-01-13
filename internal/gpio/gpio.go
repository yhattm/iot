package gpio

import "github.com/stianeikeland/go-rpio/v4"

func Init() error {
	return rpio.Open()
}

func Release() error {
	return rpio.Close()
}