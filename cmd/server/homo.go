package main

import (
	"github.com/countstarlight/homo/module/baidu"
	"github.com/sirupsen/logrus"
)

func main() {
	// Set logrus format
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "15:04:05"
	// Show colorful on windows
	customFormatter.ForceColors = true
	logrus.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true

	logrus.Infof("Homo v0.0.1")
	err := baidu.TextToSpeech("测试")
	if err != nil {
		logrus.Warnf("baidu tts failed: %s", err.Error())
	}
}
