//
// Copyright (C) 2019 Codist. - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Codist <i@codist.me>, June 2019
//

package sphinx

import (
	"encoding/binary"
	"github.com/countstarlight/homo/module/com"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"io"
	"os"
)

func Pcm2Wav(file string) error {
	// Read raw PCM data from input file.
	in, err := os.Open(file)
	if err != nil {
		return err
	}

	// Output file.
	out, err := os.Create(OutputWav)
	if err != nil {
		return err
	}
	defer com.IOClose("Save file to "+OutputWav, out)

	// 16000 Hz, 16 bit, 1 channel, WAV.
	e := wav.NewEncoder(out, sampleRate, 16, 1, 1)

	// Create new audio.IntBuffer.
	audioBuf, err := newAudioIntBuffer(in)
	if err != nil {
		return err
	}

	// Write buffer to output file. This writes a RIFF header and the PCM chunks from the audio.IntBuffer.
	if err := e.Write(audioBuf); err != nil {
		return err
	}
	if err := e.Close(); err != nil {
		return err
	}
	return nil
}

func newAudioIntBuffer(r io.Reader) (*audio.IntBuffer, error) {
	buf := audio.IntBuffer{
		Format: &audio.Format{
			NumChannels: 1,
			SampleRate:  sampleRate,
		},
	}
	for {
		var sample int16
		err := binary.Read(r, binary.LittleEndian, &sample)
		switch {
		case err == io.EOF:
			return &buf, nil
		case err != nil:
			return nil, err
		}
		buf.Data = append(buf.Data, int(sample))
	}
}
