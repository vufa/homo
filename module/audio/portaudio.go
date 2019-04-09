//
// Copyright (C) 2019 Codist. - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Codist <i@codist.me>, March 2019
//

package audio

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xlab/portaudio-go/portaudio"
)

func init() {
	// Initial PortAudio
	if err := portaudio.Initialize(); PaError(err) {
		logrus.Fatalf("PortAudio init failed: %s", PaErrorText(err))
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
