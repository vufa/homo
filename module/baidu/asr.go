//
// Copyright (C) 2019 Codist. - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Codist <i@codist.me>, June 2019
//

package baidu

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/countstarlight/homo/module/com"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

const ASR_URL = "http://vop.baidu.com/server_api"

//语音识别响应信息
//http://ai.baidu.com/docs#/ASR-API-PRO/top
type ASRResponse struct {
	CorpusNo string   `json:"corpus_no"`
	ErrMsg   string   `json:"err_msg"`
	ErrNo    int      `json:"err_no"`
	Result   []string `json:"result"`
	SN       string   `json:"sn"`
}

//语音识别参数
type ASRParams struct {
	Format   string `json:"format"`  //语音的格式，pcm 或者 wav 或者 amr。不区分大小写
	Rate     int    `json:"rate"`    //采样率，支持 8000 或者 16000
	Channel  int    `json:"channel"` //声道数，仅支持单声道，请填写固定值 1
	Cuid     string `json:"cuid"`    //用户唯一标识，用来区分用户，计算UV值。建议填写能区分用户的机器 MAC 地址或 IMEI 码，长度为60字符以内
	Token    string `json:"token"`   //开放平台获取到的开发者access_token
	Language string `json:"lan"`     //语种选择，默认中文（zh）。 中文=zh、粤语=ct、英文=en，不区分大小写
	Speech   string `json:"speech"`  //真实的语音数据 ，需要进行base64 编码。与len参数连一起使用
	Length   int    `json:"len"`     //原始语音长度，单位字节
}

type ASRParam func(params *ASRParams)

func Format(fmt string) ASRParam {

	if fmt != "pcm" && fmt != "wav" && fmt != "amr" {
		fmt = "pcm"
	}
	return func(params *ASRParams) {
		params.Format = fmt
	}
}

func Rate(rate int) ASRParam {
	if rate != 8000 && rate != 16000 {
		rate = 8000
	}
	return func(params *ASRParams) {
		params.Rate = rate
	}
}

func Channel(c int) ASRParam {
	return func(params *ASRParams) {
		params.Channel = 1 //固定值1
	}
}

func Language(lang string) ASRParam {
	if lang != "zh" && lang != "ct" && lang != "en" {
		lang = "zh"
	}
	return func(params *ASRParams) {
		params.Language = lang
	}
}

// SpeechToText 语音识别，将语音翻译成文字
func (vc *VoiceClient) SpeechToText(reader io.Reader, params ...ASRParam) ([]string, error) {

	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	if len(content) > 10*MB {
		return nil, fmt.Errorf("文件大小不能超过10M")
	}

	if err := vc.Auth(); err != nil {
		return nil, err
	}

	spch := base64.StdEncoding.EncodeToString(content)

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

	asrParams := &ASRParams{
		Format:   "pcm",
		Rate:     8000,
		Channel:  1,
		Cuid:     cuid,
		Token:    vc.AccessToken,
		Language: "zh",
		Speech:   spch,
		Length:   len(content),
	}

	for _, fn := range params {
		fn(asrParams)
	}

	asrParamsBuffer := new(bytes.Buffer)
	err = json.NewEncoder(asrParamsBuffer).Encode(asrParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", ASR_URL, asrParamsBuffer)
	if err != nil {
		return nil, err
	}
	//defer com.IOClose("Post baidu ASR api", req.Body)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer com.IOClose("Post baidu ASR api", resp.Body)

	var asrResponse *ASRResponse

	err = json.NewDecoder(resp.Body).Decode(&asrResponse)
	if err != nil {
		return nil, err
	}

	if asrResponse.ErrNo != 0 {
		if asrResponse.ErrMsg == "speech quality error." {
			return nil, ErrSpeechQuality{ErrNo: asrResponse.ErrNo, ErrMsg: asrResponse.ErrMsg}
		} else {
			return nil, fmt.Errorf("调用服务失败：%s", asrResponse.ErrMsg)
		}
	}

	return asrResponse.Result, nil

}

// Voice Recognition
// ATTENTION: the .wav file must be 8k or 16k rate with single(mono) channel.
// FYI: you can use QuickTime to record voice and Fission converting to .wav
func SpeechToText(file, format string, sampleRate int) ([]string, error) {
	client := NewVoiceClient(APIKEY, APISECRET)
	if err := client.Auth(); err != nil {
		return nil, err
	}

	f, err := os.OpenFile(file, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}

	rs, err := client.SpeechToText(
		f,
		Format(format),
		Channel(1),
		Rate(sampleRate),
	)
	if err != nil {
		return nil, err
	}
	return rs, nil
}
