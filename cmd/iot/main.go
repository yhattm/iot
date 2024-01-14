package main

import (
	"context"
	"iot/internal/gpio"

	"github.com/gookit/goutil/cflag"
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/dump"
)

var optsGpio = struct {
	pin int
	action string
}{}

func main() {
	_ = context.Background()
	app := cflag.NewApp()
	app.Desc = "iot"
	app.Version = "7.1.14"

	err := gpio.Init()
	if err == nil {
		defer gpio.Release()
	}

	gpioCmd := cflag.NewCmd("gpio", "gpio cmd")
	gpioCmd.OnAdd = func(c *cflag.Cmd) {
		c.IntVar(&optsGpio.pin, "pin", 10, "gpio pin;true;p")
		c.StringVar(&optsGpio.action, "action", "toggle", "gpio action;true;a")
	}
	gpioCmd.Func = func(c *cflag.Cmd) error {
		dump.P(optsGpio, c.Args())
		cliutil.Magentaln("gpio cmd:")
		cliutil.Infoln("gpio pin:", optsGpio.pin)
		action := gpio.Action(optsGpio.action)
		gpio.SetAction(optsGpio.pin, action)
		return nil
	}
	app.Add(gpioCmd)

	app.Run()
}
