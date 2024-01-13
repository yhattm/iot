package main

import (
	"context"

	"github.com/gookit/goutil/cflag"
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/dump"
)

var optsGpio = struct {
	pin int
}{}

func main() {
	_ = context.Background()
	app := cflag.NewApp()
	app.Desc = "iot"
	app.Version = "7.1.14"

	gpioCmd := cflag.NewCmd("gpio", "gpio cmd")
	gpioCmd.OnAdd = func(c *cflag.Cmd) {
		c.IntVar(&optsGpio.pin, "pin", 10, "gpio pin;true;t")
	}
	gpioCmd.Func = func(c *cflag.Cmd) error {
		dump.P(optsGpio, c.Args())
		cliutil.Magentaln("gpio cmd:")
		cliutil.Infoln("gpio pin:", optsGpio.pin)
		return nil
	}
	app.Add(gpioCmd)

	app.Run()
}
