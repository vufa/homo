//
// Copyright (C) 2019 Codist. - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Codist <i@codist.me>, March 2019
//

package main

import (
	"fmt"
	"github.com/countstarlight/homo/cmd/interact/config"
	"github.com/countstarlight/homo/module/baidu"
	"github.com/countstarlight/homo/module/nlu"
	"github.com/marcusolsson/tui-go"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"sync"
	"time"
)

var flags = []cli.Flag{
	cli.BoolFlag{
		EnvVar: "HOMO_INTERACT_DEBUG",
		Name:   "debug, d",
		Usage:  "start homo in debug mode",
	},
	cli.BoolFlag{
		EnvVar: "HOMO_INTENT_ONLY",
		Name:   "intent, i",
		Usage:  "start homo in intent only mode",
	},
}

func interact(ctx *cli.Context) error {
	// Set logrus format
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "15:04:05"
	// Show colorful on windows
	customFormatter.ForceColors = true
	logrus.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true

	logrus.Infof("Homo Interact v0.0.1")
	if ctx.Bool("debug") {
		logrus.Infof("Running in debug mode")
		config.DebugMode = true
	} else if ctx.Bool("intent") {
		logrus.Infof("Running homo in intent only mode")
		config.IntentOnlyMode = true
	}
	config.NewContext()
	//
	//tui begin
	//
	sidebar := tui.NewVBox(
		tui.NewLabel("Homo Interact v0.0.1"),
		tui.NewLabel("Homo 交互模式"),
		tui.NewSpacer(),
		tui.NewLabel("语音合成[开启]"),
		tui.NewLabel("语音合成API[百度]"),
		tui.NewLabel("语音自动播放[开启]"),
	)
	if ctx.Bool("debug") {
		sidebar.Append(
			tui.NewLabel("DEBUG MODE"),
		)
	} else if ctx.Bool("intent") {
		sidebar.Append(
			tui.NewLabel("INTENT-ONLY MODE"),
		)
	}
	sidebar.Append(
		tui.NewLabel("按Esc退出"),
	)
	sidebar.SetBorder(true)

	history := tui.NewVBox()

	historyScroll := tui.NewScrollArea(history)
	//cannot scroll after
	historyScroll.SetAutoscrollToBottom(true)

	historyBox := tui.NewVBox(historyScroll)
	historyBox.SetTitle("消息")
	historyBox.SetBorder(true)

	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	inputBox := tui.NewHBox(input)
	inputBox.SetTitle("输入 按Enter提交")
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	chat := tui.NewVBox(historyBox, inputBox)
	chat.SetSizePolicy(tui.Expanding, tui.Expanding)

	input.OnSubmit(func(e *tui.Entry) {
		if len(e.Text()) > 0 {
			history.Append(tui.NewHBox(
				tui.NewLabel(time.Now().Format("15:04")),
				tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("%s:", "你"))),
				tui.NewLabel(e.Text()),
				tui.NewSpacer(),
			))
			//Send text to homo core
			var (
				reply string
				err   error
			)
			playMutex := &sync.Mutex{}
			if config.IntentOnlyMode {
				replyMessage, err := nlu.ActionLocal(e.Text())
				if err != nil {
					reply = "错误: " + err.Error()
				} else {
					//display rank
					history.Append(tui.NewHBox(
						tui.NewLabel(time.Now().Format("15:04")),
						tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("%s:", "homo"))),
						tui.NewLabel(replyMessage[0]),
						tui.NewSpacer(),
					))
					//Play voice
					go func() {
						playMutex.Lock()
						er := baidu.TextToSpeech(replyMessage[0])
						playMutex.Unlock()
						if er != nil {
							rep := "语音合成出错: " + er.Error()
							history.Append(tui.NewHBox(
								tui.NewLabel(time.Now().Format("15:04")),
								tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("%s:", "homo"))),
								tui.NewLabel(rep),
								tui.NewSpacer(),
							))
						}
					}()
					//display entities
					history.Append(tui.NewHBox(
						tui.NewLabel(time.Now().Format("15:04")),
						tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("%s:", "homo"))),
						tui.NewLabel(replyMessage[1]),
						tui.NewSpacer(),
					))
					//Play voice
					go func() {
						playMutex.Lock()
						er := baidu.TextToSpeech(replyMessage[1])
						playMutex.Unlock()
						if er != nil {
							rep := "语音合成出错: " + er.Error()
							history.Append(tui.NewHBox(
								tui.NewLabel(time.Now().Format("15:04")),
								tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("%s:", "homo"))),
								tui.NewLabel(rep),
								tui.NewSpacer(),
							))
						}
					}()
					reply = replyMessage[2]
				}
			} else {
				reply, err = nlu.ChatWithCore(e.Text())
				if err != nil {
					reply = "连接到Homo Core出错: " + err.Error()
				}
			}
			history.Append(tui.NewHBox(
				tui.NewLabel(time.Now().Format("15:04")),
				tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("%s:", "homo"))),
				tui.NewLabel(reply),
				tui.NewSpacer(),
			))
			//Play voice
			go func() {
				playMutex.Lock()
				er := baidu.TextToSpeech(reply)
				playMutex.Unlock()
				if er != nil {
					rep := "语音合成出错: " + er.Error()
					history.Append(tui.NewHBox(
						tui.NewLabel(time.Now().Format("15:04")),
						tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("%s:", "homo"))),
						tui.NewLabel(rep),
						tui.NewSpacer(),
					))
				}
			}()
			input.SetText("")
		}
	})

	root := tui.NewHBox(sidebar, chat)

	ui, err := tui.New(root)
	if err != nil {
		logrus.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	//ui.SetKeybinding("Up", func() { historyScroll.Scroll(0, -1) })
	//ui.SetKeybinding("Down", func() { historyScroll.Scroll(0, 1) })

	if err := ui.Run(); err != nil {
		logrus.Fatal(err)
	}
	return nil
}

func before(c *cli.Context) error { return nil }
