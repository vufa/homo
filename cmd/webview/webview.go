//
// Copyright (C) 2019 Codist. - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Codist <i@codist.me>, April 2019
//

package main

//go:generate go-bindata -pkg $GOPACKAGE -o bindata.go -prefix assets/ assets/...

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime"
	"net"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/zserge/webview"
)

var flags = []cli.Flag{
	cli.BoolFlag{
		EnvVar: "HOMO_WEBVIEW_DEBUG",
		Name:   "debug, d",
		Usage:  "start homo webview in debug mode",
	},
}

func startServer() string {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer ln.Close()
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			if len(path) > 0 && path[0] == '/' {
				path = path[1:]
			}
			if path == "" {
				path = "index.html"
			}
			//fmt.Printf("path %s\n", path)
			if bs, err := Asset(path); err != nil {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.Header().Add("Content-Type", mime.TypeByExtension(filepath.Ext(path)))
				_, err = io.Copy(w, bytes.NewBuffer(bs))
				if err != nil {
					panic(err)
				}
			}
		})
		log.Fatal(http.Serve(ln, nil))
	}()
	return "http://" + ln.Addr().String()
}

// Task is a data model type, it contains information about task name and status (done/not done).
type Task struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

// Tasks is a global data model, to keep things simple.
var Tasks = []Task{}

func render(w webview.WebView, tasks []Task) {
	b, err := json.Marshal(tasks)
	if err == nil {
		err = w.Eval(fmt.Sprintf("rpc.render(%s)", string(b)))
		if err != nil {
			panic(err)
		}
	}
}

func handleRPC(w webview.WebView, data string) {
	switch {
	case strings.HasPrefix(data, "message:"):
		fmt.Printf("发送的消息: %s\n", strings.TrimPrefix(data, "message:"))
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
	url := startServer()
	if ctx.Bool("debug") {
		logrus.Infof("Running in debug mode")
	}

	w := webview.New(webview.Settings{
		Width:                  400,
		Height:                 600,
		Title:                  AppName,
		URL:                    url,
		Debug:                  ctx.Bool("debug"),
		ExternalInvokeCallback: handleRPC,
	})
	defer w.Exit()
	w.Run()
}

func before(c *cli.Context) error { return nil }
