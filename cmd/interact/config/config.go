//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, March 2019
//

package config

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"sync"
)

var (
	BeepSpeakerInited bool
	DebugMode         bool
	IntentOnlyMode    bool
	VoicePlayMutex    sync.Mutex
)

func init() {
	BeepSpeakerInited = false
	DebugMode = false
	IntentOnlyMode = false
}

func NewContext() {
	//Initial portaudio
	/*err := portaudio.Initialize()
	if err != nil {
		logrus.Fatalf("Initial portaudio failed: %s", err.Error())
	}*/
}

func Terminal(c *cli.Context) error {
	logrus.Infof("退出，开始结束进程...")
	/*err := portaudio.Terminate()
	if err != nil {
		logrus.Warnf("Close portaudio failed", err.Error())
		return err
	}*/
	return nil
}
