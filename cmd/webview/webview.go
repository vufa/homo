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
	"github.com/countstarlight/homo/cmd/webview/config"
	"github.com/countstarlight/homo/module/baidu"
	"github.com/countstarlight/homo/module/com"
	"github.com/countstarlight/homo/module/nlu"
	"github.com/sirupsen/logrus"
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

func startServer() string {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer com.IOClose("webview ln", ln)
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
	w.Dispatch(func() {
		err := w.Eval("chatWindow.think()")
		if err != nil {
			logrus.Warning("botTypingAnimate w.Eval failed: %s", err.Error())
		}
	})
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
	w.Dispatch(func() {
		err = w.Eval(fmt.Sprintf("chatWindow.talk(%s, \"message\")", string(b)))
		if err != nil {
			logrus.Warning("sendReply: w.Eval failed: %s", err.Error())
		}
	})
}

func handleRPC(w webview.WebView, data string) {
	switch {
	case strings.HasPrefix(data, "message:"):
		msg := strings.TrimPrefix(data, "message:")
		//fmt.Printf("发送的消息: %s\n", msg)
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
				config.VoicePlayMutex.Lock()
				err = baidu.TextToSpeech(sent)
				config.VoicePlayMutex.Unlock()
				if err != nil {
					w.Dispatch(func() {
						sendReply(w, []string{"语音合成出错: " + err.Error()})
					})
				}
			}
		}()
	}
}
