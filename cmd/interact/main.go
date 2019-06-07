//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, March 2019
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
