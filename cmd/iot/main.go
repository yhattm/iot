package main

import (
	"context"
	"iot/internal/gpio"
	"net/http"

	"github.com/gookit/goutil/cflag"
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/dump"

	"github.com/gin-gonic/gin"
)

const Version = "7.1.14"

var optsGpio = struct {
	pin    int
	action string
}{}

func main() {
	_ = context.Background()
	app := cflag.NewApp()
	app.Desc = "iot"
	app.Version = Version

	err := gpio.Init()
	if err == nil {
		defer gpio.Release()
	}

	gpioCmd := cflag.NewCmd("gpio", "gpio cmd")
	gpioCmd.OnAdd = func(c *cflag.Cmd) {
		c.IntVar(&optsGpio.pin, "pin", 14, "gpio pin;true;p")
		c.StringVar(&optsGpio.action, "action", "open", "gpio action;true;a")
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

	httpCmd := cflag.NewCmd("http", "http service")
	httpCmd.Func = func(c *cflag.Cmd) error {
		server := gin.Default()
		server.GET("/", handleVersion)
		server.GET("/open", handleOpen)
		return server.Run(":8080")
	}
	app.Add(httpCmd)

	app.Run()
}

func handleOpen(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.Set("content-type", "application/json")
	gpio.SetAction(14, gpio.ActionLowToHigh)
	c.Status(http.StatusOK)
	c.Writer.Write([]byte("opened"))
	c.Writer.Flush()
}

func handleVersion(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.Set("content-type", "application/json")
	c.Status(http.StatusOK)
	c.Writer.Write([]byte(Version))
	c.Writer.Flush()
}
