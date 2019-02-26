package baidu

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/bobertlo/go-mpg123/mpg123"
	"github.com/countstarlight/homo/module/com"
	"github.com/gordonklaus/portaudio"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"

	"net"
)

const TTS_URL = "http://tsn.baidu.com/text2audio"

//TextToSpeech 对接baidu tts rest api
//https://ai.baidu.com/docs#/TTS-API/top
func (vc *VoiceClient) TextToSpeech(txt string) ([]byte, error) {

	if len(txt) >= 1024 {
		return nil, fmt.Errorf("Input text too long: %d > 1024", len(txt))
	}
	if err := vc.Auth(); err != nil {
		return nil, err
	}

	var cuid string
	netitfs, err := net.Interfaces()
	if err != nil {
		cuid = "anonymous"
	} else {
		for _, itf := range netitfs {
			if cuid = itf.HardwareAddr.String(); len(cuid) > 0 {
				break
			}
		}
	}

	resp, err := http.PostForm(TTS_URL, url.Values{
		"tex":  {txt},
		"tok":  {vc.AccessToken},
		"cuid": {cuid},
		"ctp":  {"1"},
		"lan":  {"zh"},
		"spd":  {"5"},
		"pit":  {"5"},
		"vol":  {"5"},
		"per":  {"0"},
		"aue":  {"3"}, //mp3 format
	})
	if err != nil {
		return nil, err
	}
	defer com.IOClose("Post baidu tts api", resp.Body)

	//通过Content-Type的头部来确定是否服务端合成成功。
	//http://ai.baidu.com/docs#/TTS-API/top
	respHeader := resp.Header
	contentType, ok := respHeader["Content-Type"]
	if !ok {
		return nil, fmt.Errorf("No Content-Type Set.")
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if contentType[0] == "audio/mp3" {
		return respBody, nil
	} else {
		return nil, fmt.Errorf("调用服务失败：%s", string(respBody))
	}

}

const (
	// This Api Key and Api Secret is just for example,
	// you should get your own first.
	APIKEY    = "MDNsII2jkUtbF729GQOZt7FS"
	APISECRET = "0vWCVCLsbWHMSH1wjvxaDq4VmvCZM2O9"
)

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
)

type VoiceClient struct {
	*Client
}

func NewVoiceClient(apiKey, apiSecret string) *VoiceClient {
	return &VoiceClient{
		Client: NewClient(apiKey, apiSecret),
	}
}

// Voice Composition
func TextToSpeech(text string) error {
	client := NewVoiceClient(APIKEY, APISECRET)
	voiceData, err := client.TextToSpeech(text)
	if err != nil {
		return err
	}

	//Remove previous file
	if com.IsFile("tmp/tts/tmp.mp3") {
		err = os.Remove("tmp/tts/tmp.mp3")
		if err != nil {
			return err
		}
	}

	f, err := os.OpenFile("tmp/tts/tmp.mp3", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer com.IOClose("Save baidu tts to file", f)
	if _, err := f.Write(voiceData); err != nil {
		return err
	}

	//
	//decode mp3 voice data
	//
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	decoder, err := mpg123.NewDecoder("")
	err = decoder.Open("tmp/tts/tmp.mp3")
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
