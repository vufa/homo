//
// Copyright (C) 2019 Codist. - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Codist <i@codist.me>, June 2019
//

package sphinx

import (
	"github.com/countstarlight/homo/cmd/webview/config"
	"github.com/countstarlight/homo/module/audio"
	"github.com/sirupsen/logrus"
	"github.com/xlab/pocketsphinx-go/sphinx"
	"github.com/xlab/portaudio-go/portaudio"
)

func LoadCMUSphinxCapture() {
	config.WakeUpWait.Add(1)
	// Init CMUSphinx
	cfg := sphinx.NewConfig(
		sphinx.HMMDirOption("sphinx/en-us/en-us"),
		sphinx.DictFileOption("sphinx/homo/homo.dic"),
		sphinx.LMFileOption("sphinx/homo/homo.lm.bin"),
		sphinx.SampleRateOption(sampleRate),
	)
	//Specify output dir for RAW recorded sound files (s16le). Directory must exist.
	sphinx.RawLogDirOption("tmp/record")(cfg)

	sphinx.LogFileOption("log/sphinx.log")(cfg)

	logrus.Info("开始加载 CMU PhocketSphinx...")
	logrus.Info("开始加载唤醒模型...")
	dec, err := sphinx.NewDecoder(cfg)
	if err != nil {
		logrus.Fatalf("sphinx.NewDecoder failed: %s", err.Error())
	}
	defer dec.Destroy()

	dec.SetRawDataSize(300000)

	l := &Listener{
		dec: dec,
	}

	var stream *portaudio.Stream
	errStr := portaudio.OpenDefaultStream(&stream, channels, 0, sampleFormat, sampleRate, samplesPerChannel, l.paCallback, nil)
	if audio.PaError(errStr) {
		logrus.Fatalf("PortAudio failed: %s", audio.PaErrorText(errStr))
	}
	defer func() {
		if errStr := portaudio.CloseStream(stream); audio.PaError(errStr) {
			logrus.Warnf("PortAudio error:", audio.PaErrorText(errStr))
		}
	}()
	if errStr := portaudio.StartStream(stream); audio.PaError(errStr) {
		logrus.Fatalf("PortAudio error: %s", audio.PaErrorText(errStr))
	}
	defer func() {
		if errStr := portaudio.StopStream(stream); audio.PaError(errStr) {
			logrus.Fatalf("PortAudio error:", audio.PaErrorText(errStr))
		}
	}()
	if !dec.StartUtt() {
		logrus.Fatalln("Sphinx failed to start utterance")
	}
	logrus.Infof("开始从麦克风检测唤醒词：采样率[%dHz] 通道数[%d]", sampleRate, channels)
	config.WakeUpWait.Wait()
	config.WakeUpd = true
}
