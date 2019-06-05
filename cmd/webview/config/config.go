//
// Copyright (C) 2019 Codist. - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Codist <i@codist.me>, March 2019
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
	DebugMode      bool
	VoicePlayMutex sync.Mutex
	// Make sphinx keep capturing audio input
	SphinxLoop sync.WaitGroup
	WakeUpWait sync.WaitGroup

	//Auto convert raw pcm buffer to wav
	RawToWav bool
	//WebViewWait     sync.WaitGroup
)

func init() {
	WakeUpd = false
	DebugMode = false
	RawToWav = false
}

func Terminal(c *cli.Context) error {
	logrus.Infof("退出，开始结束PortAudio...")
	return audio.PaTerminate()
}
