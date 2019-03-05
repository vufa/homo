package audio

import (
	"bytes"
	"encoding/binary"
	"github.com/bobertlo/go-mpg123/mpg123"
	"github.com/countstarlight/homo/module/com"
	"github.com/gordonklaus/portaudio"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func PortAudioPlayMp3(fileName string) error {
	//
	//decode mp3 voice data
	//
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	decoder, err := mpg123.NewDecoder("")
	err = decoder.Open(fileName)
	if err != nil {
		return err
	}
	defer com.IOClose("mpg123 decoder", decoder)
	// get audio format information
	rate, channels, _ := decoder.GetFormat()

	// make sure output format does not change
	decoder.FormatNone()
	decoder.Format(rate, channels, mpg123.ENC_SIGNED_16)

	out := make([]int16, 8192)
	stream, err := portaudio.OpenDefaultStream(0, channels, float64(rate), len(out), &out)
	if err != nil {
		return err
	}
	defer com.IOClose("portaudio stream", stream)
	err = stream.Start()
	if err != nil {
		return err
	}
	defer func() {
		err = stream.Stop()
		if err != nil {
			logrus.Warnf("Close portaudio stream failed: %s", err.Error())
		}
	}()

	for {
		audio := make([]byte, 2*len(out))
		_, err = decoder.Read(audio)
		if err == mpg123.EOF {
			break
		}
		if err != nil {
			return err
		}

		err = binary.Read(bytes.NewBuffer(audio), binary.LittleEndian, out)
		if err != nil {
			return err
		}
		err = stream.Write()
		if err != nil {
			return err
		}
		select {
		case <-sig:
			return nil
		default:
		}
	}
	return nil
}
