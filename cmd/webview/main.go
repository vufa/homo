//
// Copyright (C) 2019 Codist. - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Codist <i@codist.me>, March 2019
//

package main

import (
	"fmt"
	"github.com/countstarlight/homo/cmd/webview/config"
	"github.com/countstarlight/homo/module/sphinx"
	"github.com/countstarlight/homo/module/view"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//setting.MaxConcurrency = runtime.NumCPU() * 2

	// initialize global pseudo random generator
	rand.Seed(time.Now().Unix())
}

var flags = []cli.Flag{
	cli.BoolFlag{
		EnvVar: "HOMO_WEBVIEW_DEBUG",
		Name:   "debug, d",
		Usage:  "start homo webview in debug mode",
	},
}

// Greeting list
var Greetings = [...]string{"我在听，请说", "Hi，有什么我可以帮你的吗？"}

func main() {
	app := cli.NewApp()
	app.Name = config.AppName
	app.Version = config.AppVersion
	app.Usage = "Help"
	app.Action = lanchWebview
	app.Flags = flags
	app.Before = before
	app.After = config.Terminal
	if err := app.Run(os.Args); err != nil {
		logrus.Fatalf("[homo-webview]%s", err.Error())
	}
}

func lanchWebview(ctx *cli.Context) {
	if ctx.Bool("debug") {
		config.DebugMode = true
		// Set logrus format
		// Print file name and line code
		logrus.SetReportCaller(true)
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "15:04:05",
			// Show colorful on windows
			ForceColors:   true,
			FullTimestamp: true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				repopath := fmt.Sprintf("%s/src/github.com/countstarlight/homo/", os.Getenv("GOPATH"))
				filename := strings.Replace(f.File, repopath, "", -1)
				r := strings.Split(f.Function, ".")
				return fmt.Sprintf("%s()", r[len(r)-1]), fmt.Sprintf("%s:%d", filename, f.Line)
			},
		})
		logrus.Infof("Running in debug mode")
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "15:04:05",
			// Show colorful on windows
			ForceColors:   true,
			FullTimestamp: true,
		})
	}
	// Init webview
	view.InitWebView(config.AppName, config.DebugMode)
	//defer w.Exit()
	//
	// Prepare wake up function
	//
	config.WakeUpWait.Add(1)
	go sphinx.LoadCMUSphinx()
	config.WakeUpWait.Wait()

	logrus.Infof("唤醒成功，开始唤起界面...")

	go func() {
		Greeting := Greetings[rand.Intn(len(Greetings))]
		view.SendReplyWithVoice([]string{Greeting})
	}()

	view.Run()
}

func before(c *cli.Context) error { return nil }
