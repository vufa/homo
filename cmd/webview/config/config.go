//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, March 2019
//

package config

import (
	"github.com/countstarlight/homo/module/audio"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"sync"
)

var (
	WakeUpd        bool
	WakeUpDisabled bool
	DebugMode      bool
	OfflineMode    bool
	// display intent and entities
	AnalyticalMode bool
	VoicePlayMutex sync.Mutex
	// Make sphinx keep capturing audio input
	SphinxLoop sync.WaitGroup
	WakeUpWait sync.WaitGroup

	//Auto convert raw pcm buffer to wav
	RawToWav bool
	//WebViewWait     sync.WaitGroup
)

const (
	AppName    = "Homo Webview"
	AppVersion = "0.0.1"
)

func init() {
	WakeUpd = false
	DebugMode = false
	RawToWav = false
	AnalyticalMode = false
}

func Terminal(c *cli.Context) error {
	logrus.Infof("退出，开始结束PortAudio...")
	return audio.PaTerminate()
}
