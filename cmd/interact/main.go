//
// Copyright (C) 2019 Codist. - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Codist <i@codist.me>, March 2019
//

package main

import (
	"github.com/countstarlight/homo/cmd/interact/config"
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
	app.After = config.Terminal
	if err := app.Run(os.Args); err != nil {
		logrus.Fatalf("[homo-interact]%s", err.Error())
	}
}
