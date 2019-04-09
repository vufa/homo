//
// Copyright (C) 2019 Codist. - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Codist <i@codist.me>, March 2019
//

package main

import (
	"github.com/countstarlight/homo/cmd/webview/config"
	"github.com/countstarlight/homo/module/baidu"
	"github.com/countstarlight/homo/module/wakeup"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/zserge/webview"
	"os"
	"runtime"
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

	// Set logrus format
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "15:04:05"
	// Show colorful on windows
	customFormatter.ForceColors = true
	logrus.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true
	if ctx.Bool("debug") {
		logrus.Infof("Running in debug mode")
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
		sendReply(w, []string{"我在听，请说"})
		time.Sleep(time.Second)
		config.VoicePlayMutex.Lock()
		err := baidu.TextToSpeech("我在听，请说")
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
