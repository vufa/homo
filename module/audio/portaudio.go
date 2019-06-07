//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, March 2019
//

package audio

import (
	"fmt"
	"github.com/countstarlight/homo/module/com"
	"github.com/sirupsen/logrus"
	"github.com/xlab/portaudio-go/portaudio"
)

func init() {
	// Initial PortAudio
	if _, err := com.CaptureWithCGo(func() {
		if err := portaudio.Initialize(); PaError(err) {
			//logrus.Fatalf("PortAudio init failed: %s", PaErrorText(err))
			//logrus print will be captured
			panic(fmt.Errorf("PortAudio init failed: %s", PaErrorText(err)))
		}
	}); err != nil {
		logrus.Fatalf("Capture PortAudio initial output failed: %s", err.Error())
	}

}
func PaError(err portaudio.Error) bool {
	return portaudio.ErrorCode(err) != portaudio.PaNoError
}

func PaErrorText(err portaudio.Error) string {
	return portaudio.GetErrorText(err)
}

func PaTerminate() error {
	if err := portaudio.Terminate(); PaError(err) {
		return fmt.Errorf("PortAudio term failed: %s", PaErrorText(err))
	}
	return nil
}
