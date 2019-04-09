//
// Copyright (C) 2019 Codist. - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Codist <i@codist.me>, March 2019
//

package audio

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"os"
	"time"
)

var (
	BeepSpeakerInited bool
)

func init() {
	BeepSpeakerInited = false
}
func BeepPlayMp3(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	// Decode the .mp3 File, if you have a .wav file, use wav.Decode(f)
	s, format, _ := mp3.Decode(f)

	// Init the Speaker with the SampleRate of the format and a buffer size of 1/10s
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		return err
	}

	// Channel, which will signal the end of the playback.
	playing := make(chan struct{})

	// Now we Play our Streamer on the Speaker
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		// Callback after the stream Ends
		close(playing)
	})))
	<-playing
	return nil
}

func BeepPlayWav(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	// Decode the .mp3 File, if you have a .wav file, use wav.Decode(f)
	s, format, _ := wav.Decode(f)

	// Init the Speaker with the SampleRate of the format and a buffer size of 1/10s
	if !BeepSpeakerInited {
		err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		if err != nil {
			return err
		}
		BeepSpeakerInited = true
	}

	// Channel, which will signal the end of the playback.
	playing := make(chan struct{})

	// Now we Play our Streamer on the Speaker
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		// Callback after the stream Ends
		close(playing)
	})))
	<-playing
	return nil
}
