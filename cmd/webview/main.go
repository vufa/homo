//
// Copyright (C) 2019 Codist. - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Codist <i@codist.me>, March 2019
//

package main

import (
	"github.com/countstarlight/homo/module/setting"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//setting.MaxConcurrency = runtime.NumCPU() * 2
}

const AppName = "Homo Webview"
const AppVersion = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = AppName
	app.Version = AppVersion
	app.Usage = "Help"
	app.Action = lanchWebview
	app.Flags = flags
	app.Before = before
	app.After = setting.Terminal
	if err := app.Run(os.Args); err != nil {
		logrus.Fatalf("[homo-webview]%s", err.Error())
	}
}
