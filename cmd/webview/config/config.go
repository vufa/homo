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

	// Mode
	DebugMode      bool
	SilenceMode    bool
	OfflineMode    bool
	InterruptMode  bool
	AnalyticalMode bool // display intent and entities

	// Flag
	IsPlayingVoice bool

	// Lock
	VoicePlayMutex sync.Mutex
	SphinxLoop     sync.WaitGroup // Make sphinx keep capturing audio input
	WakeUpWait     sync.WaitGroup

	RecordThreshold int

	//Auto convert raw pcm buffer to wav
	RawToWav bool
)

const (
	AppName    = "Homo Webview"
	AppVersion = "0.0.1"
)

func init() {
	// DebugMode = false
	// WakeUpDisabled = false
	RawToWav = false
	AnalyticalMode = false
	RecordThreshold = 50000
}

func Terminal(c *cli.Context) error {
	logrus.Infof("退出，开始结束PortAudio...")
	return audio.PaTerminate()
}
