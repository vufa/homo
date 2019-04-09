//
// Copyright (C) 2019 Codist. - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Codist <i@codist.me>, March 2019
//

package config

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"sync"
)

var (
	WakeUpd        bool
	DebugMode      bool
	VoicePlayMutex sync.Mutex
)

func Terminal(c *cli.Context) error {
	logrus.Infof("退出，开始结束PortAudio...")
	/*err := portaudio.Terminate()
	if err != nil {
		logrus.Warnf("Close portaudio failed", err.Error())
		return err
	}*/
	return nil
}
