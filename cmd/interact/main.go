package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//setting.MaxConcurrency = runtime.NumCPU() * 2
}

func main() {
	app := cli.NewApp()
	app.Name = "homo-interact"
	app.Version = "0.0.1"
	app.Usage = "Homo interact"
	app.Action = interact
	app.Flags = flags
	app.Before = before
	if err := app.Run(os.Args); err != nil {
		logrus.Fatalf("[homo-interact]%s", err.Error())
	}
}
