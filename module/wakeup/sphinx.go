//
// Copyright (C) 2019 Codist. - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Codist <i@codist.me>, March 2019
//

package wakeup

import (
	"github.com/countstarlight/homo/cmd/webview/config"
	"github.com/countstarlight/homo/module/audio"
	"github.com/sirupsen/logrus"
	"github.com/xlab/pocketsphinx-go/sphinx"
	"github.com/xlab/portaudio-go/portaudio"
	"unsafe"
)

const (
	samplesPerChannel = 512
	sampleRate        = 16000
	channels          = 1
	sampleFormat      = portaudio.PaInt16
)

type Listener struct {
	inSpeech   bool
	uttStarted bool
	dec        *sphinx.Decoder
}

func LoadCMUSphinx() {
	config.WakeUpWait.Add(1)
	// Init CMUSphinx
	cfg := sphinx.NewConfig(
		sphinx.HMMDirOption("data/wake_up/en-us/en-us"),
		sphinx.DictFileOption("data/wake_up/homo/homo.dic"),
		sphinx.LMFileOption("data/wake_up/homo/homo.lm.bin"),
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
	config.WakeUpWait.Wait()
}

// paCallback: for simplicity reasons we process raw audio with sphinx in the this stream callback,
// never do that for any serious applications, use a buffered channel instead.
func (l *Listener) paCallback(input unsafe.Pointer, _ unsafe.Pointer, sampleCount uint,
	_ *portaudio.StreamCallbackTimeInfo, _ portaudio.StreamCallbackFlags, _ unsafe.Pointer) int32 {

	const (
		statusContinue = int32(portaudio.PaContinue)
		statusAbort    = int32(portaudio.PaAbort)
	)

	in := (*(*[1 << 24]int16)(input))[:int(sampleCount)*channels]
	// ProcessRaw with disabled search because callback needs to be relatime
	_, ok := l.dec.ProcessRaw(in, true, false)
	// log.Printf("processed: %d frames, ok: %v", frames, ok)
	if !ok {
		return statusAbort
	}
	if l.dec.IsInSpeech() {
		l.inSpeech = true
		if !l.uttStarted {
			l.uttStarted = true
			logrus.Info("检测唤醒词...")
		}
	} else if l.uttStarted {
		// speech -> silence transition, time to start new utterance
		l.dec.EndUtt()
		l.uttStarted = false
		l.report() // report results
		if !l.dec.StartUtt() {
			logrus.Fatalln("Sphinx failed to start utterance")
		}
	}
	return statusContinue
}

func (l *Listener) report() {
	hyp, _ := l.dec.Hypothesis()
	if len(hyp) > 0 {
		//logrus.Printf("    > hypothesis: %s", hyp)
		if hyp == "homo" || hyp == "como" {
			logrus.Info("检测到唤醒词，开始唤醒")
			config.WakeUpWait.Done()
		}
		return
	}
	logrus.Println("没有检测到唤醒词")
}
