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
	"github.com/countstarlight/homo/module/baidu"
	"github.com/countstarlight/homo/module/wakeup"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/zserge/webview"
	"os"
	"runtime"
	"strings"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//setting.MaxConcurrency = runtime.NumCPU() * 2
}

const AppName = "Homo Webview"
const AppVersion = "0.0.1"

var flags = []cli.Flag{
	cli.BoolFlag{
		EnvVar: "HOMO_WEBVIEW_DEBUG",
		Name:   "debug, d",
		Usage:  "start homo webview in debug mode",
	},
}

const Greeting = "我在听，请说"

func main() {
	app := cli.NewApp()
	app.Name = AppName
	app.Version = AppVersion
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
			ForceColors: true,
		})
	}
	w := webview.New(webview.Settings{
		Width:                  900,
		Height:                 700,
		Resizable:              true,
		Title:                  AppName,
		URL:                    startServer(),
		Debug:                  ctx.Bool("debug"),
		ExternalInvokeCallback: handleRPC,
	})
	defer w.Exit()
	//
	// Prepare wake up function
	//
	wakeup.LoadCMUSphinx()
	go func() {
		sendReply(w, []string{Greeting})
		time.Sleep(time.Second)
		config.VoicePlayMutex.Lock()
		err := baidu.TextToSpeech(Greeting)
		config.VoicePlayMutex.Unlock()
		if err != nil {
			w.Dispatch(func() {
				sendReply(w, []string{"语音合成出错: " + err.Error()})
			})
		}
	}()

	// Run webview
	w.Run()
}

func before(c *cli.Context) error { return nil }
