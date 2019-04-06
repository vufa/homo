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
	"github.com/countstarlight/homo/module/baidu"
	"github.com/countstarlight/homo/module/nlu"
	"github.com/countstarlight/homo/module/setting"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/zserge/webview"
	"io"
	"log"
	"mime"
	"net"
	"net/http"
	"path/filepath"
	"strings"
	"time"
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

type Message struct {
	Says []string `json:"says"`
}
type HomoReply struct {
	Msg Message `json:"message"`
}

func typingAnimate(w webview.WebView) {
	err := w.Eval("chatWindow.think()")
	if err != nil {
		logrus.Warning("botTypingAnimate w.Eval failed: %s", err.Error())
	}
}
func sendReply(w webview.WebView, message []string) {
	b, err := json.Marshal(HomoReply{
		Msg: Message{
			Says: message,
		},
	})
	if err != nil {
		logrus.Warning("sendReply: json.Marshal failed: %s", err.Error())
	}
	err = w.Eval(fmt.Sprintf("chatWindow.talk(%s, \"message\")", string(b)))
	if err != nil {
		logrus.Warning("sendReply: w.Eval failed: %s", err.Error())
	}
}

func handleRPC(w webview.WebView, data string) {
	switch {
	case strings.HasPrefix(data, "message:"):
		msg := strings.TrimPrefix(data, "message:")
		//fmt.Printf("发送的消息: %s\n", msg)
		//typingAnimate(w)
		go func() {
			var reply []string
			replyMessage, err := nlu.ActionLocal(msg)
			if err != nil {
				reply = []string{"错误: " + err.Error()}
			} else {
				reply = replyMessage
			}
			w.Dispatch(func() {
				//sendReply(w, []string{"你好", "今天天气不错", "不是吗"})
				sendReply(w, reply)
			})
			//Play voice
			time.Sleep(time.Second)
			for _, sent := range replyMessage {
				setting.VoicePlayMutex.Lock()
				err = baidu.TextToSpeech(sent)
				setting.VoicePlayMutex.Unlock()
				if err != nil {
					w.Dispatch(func() {
						//sendReply(w, []string{"你好", "今天天气不错", "不是吗"})
						sendReply(w, []string{"语音合成出错: " + err.Error()})
					})
				}
			}
		}()
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
		Width:                  900,
		Height:                 700,
		Title:                  AppName,
		URL:                    url,
		Debug:                  ctx.Bool("debug"),
		ExternalInvokeCallback: handleRPC,
	})
	defer w.Exit()
	w.Run()
}

func before(c *cli.Context) error { return nil }
