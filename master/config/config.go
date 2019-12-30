//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, March 2019
//

package config

import (
	"fmt"
	"github.com/countstarlight/homo/utils/com"
	"github.com/go-ini/ini"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var (
	//App settings
	AppPath    string
	LogPath    string
	OTALogPath string

	// Mode
	DebugMode      bool
	SilenceMode    bool
	OfflineMode    bool
	InterruptMode  bool
	AnalyticalMode bool // display intent and entities

	// Global config
	Cfg         *ini.File
	ConfFile    string
	WorkDirPath string

	// PortAudio
	RawDir   string
	InputRaw string
	InputWav string

	//sphinx
	HMMDirEn        string
	DictFileEn      string
	LMFileEn        string
	SphinxLogFile   string
	RecordThreshold int

	// Nlu
	ConversationAPI string
	ParseAPI        string
	NluProject      string
	NluModel        string

	// baidu
	BaiduASRAPI         string
	BaiduTTSAPI         string
	BaiduVoiceAuthUrl   string
	BaiduVoiceAPIKey    string
	BaiduVoiceAPISecret string

	// TTS
	TTSDir     string
	TTSOutFile string

	// Flag
	IsPlayingVoice bool
	WakeUpd        bool
	WakeUpDisabled bool

	// Lock
	VoicePlayMutex sync.Mutex
	SphinxLoop     sync.WaitGroup // Make sphinx keep capturing audio input
	WakeUpWait     sync.WaitGroup

	//Auto convert raw pcm buffer to wav
	RawToWav bool
)

const (
	AppName    = "Homo"
	AppVersion = "0.0.1"
)

// execPath returns the executable path.
func execPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Abs(file)
}

func init() {
	RawToWav = false
	AnalyticalMode = false
	RecordThreshold = 50000

	var err error
	if AppPath, err = execPath(); err != nil {
		log.Fatalf("Fail to get app path: %s\n", err.Error())
	}

	// Note: we don't use path.Dir here because it does not handle case
	//	which path starts with two "/" in Windows: "//psf/Home/..."
	AppPath = strings.Replace(AppPath, "\\", "/", -1)
}

// WorkDir returns absolute path of work directory.
func WorkDir() (string, error) {
	wd := os.Getenv("HOMO_WORK_DIR")
	if len(wd) > 0 {
		return wd, nil
	}

	i := strings.LastIndex(AppPath, "/")
	if i == -1 {
		return AppPath, nil
	}
	return AppPath[:i], nil
}

func LoadConfig() error {
	var err error
	if WorkDirPath == "" {
		WorkDirPath, err = WorkDir()
		if err != nil {
			return fmt.Errorf("Get work directory failed: %s", err.Error())
		}
	}
	// ConfFile = path.Join(WorkDirPath, "conf/app.ini")
	if ConfFile == "" {
		ConfFile = path.Join(WorkDirPath, "conf/app.ini")
	}

	if !com.IsFile(ConfFile) {
		return fmt.Errorf("没有找到配置文件 %s , 如果是第一次运行，请拷贝一份 conf/example_app.ini 到 conf/app.ini", ConfFile)
	}

	Cfg, err = ini.Load(ConfFile)
	if err != nil {
		return fmt.Errorf("Parse config file %s failed: %s", ConfFile, err.Error())
	}

	Cfg.NameMapper = ini.AllCapsUnderscore

	// Load log config
	sec := Cfg.Section("log")
	LogPath = sec.Key("ROOT_PATH").MustString(path.Join(WorkDirPath, "log"))

	// Create log path
	if !com.PathExists(LogPath) {
		err := os.MkdirAll(LogPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("Create path %s failed: %s", LogPath, err.Error())
		}
	}

	// Load NLU config
	sec = Cfg.Section("nlu")
	ConversationAPI = sec.Key("CONVERSATION_API").MustString("http://localhost:5005/conversations/default/respond")
	ParseAPI = sec.Key("PARSE_API").MustString("http://localhost:5000/parse")
	NluProject = sec.Key("PROJECT").MustString("rasa")
	NluModel = sec.Key("MODEL").MustString("ini")

	// Load baidu config
	sec = Cfg.Section("baidu")
	BaiduASRAPI = sec.Key("ASR_API").MustString("http://vop.baidu.com/server_api")
	BaiduTTSAPI = sec.Key("TTS_API").MustString("http://tsn.baidu.com/text2audio")
	BaiduVoiceAuthUrl = sec.Key("VOICE_AUTH_URL").MustString("https://openapi.baidu.com/oauth/2.0/token")
	BaiduVoiceAPIKey = sec.Key("VOICE_API_KEY").MustString("MDNsII2jkUtbF729GQOZt7FS")
	BaiduVoiceAPISecret = sec.Key("VOICE_API_SECRET").MustString("0vWCVCLsbWHMSH1wjvxaDq4VmvCZM2O9")

	// Load tts config
	sec = Cfg.Section("tts")
	TTSDir = sec.Key("TTS_DIR").MustString(path.Join(WorkDirPath, "tmp/tts"))
	TTSOutFile = sec.Key("TTS_OUT_FILE").MustString(path.Join(WorkDirPath, "tmp/tts/tmp.wav"))

	// Create path
	if !com.PathExists(TTSDir) {
		err := os.MkdirAll(TTSDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("Create path %s failed: %s", TTSDir, err.Error())
		}
	}

	// Update config file
	return UpdateConfigFile()
}

func UpdateConfigFile() error {
	cfg := ini.Empty()
	if com.IsFile(ConfFile) {
		// Keeps custom settings if there is already something.
		if err := cfg.Append(ConfFile); err != nil {
			return fmt.Errorf("Fail to load conf '%s': %s", ConfFile, err.Error())
		}
	}

	// Update log config
	cfg.Section("log").Key("ROOT_PATH").SetValue(LogPath)

	// Update PortAudio config
	cfg.Section("portaudio").Key("RAW_DIR").SetValue(RawDir)
	cfg.Section("portaudio").Key("INPUT_RAW").SetValue(InputRaw)
	cfg.Section("portaudio").Key("INPUT_WAV").SetValue(InputWav)

	// Update sphinx config
	cfg.Section("sphinx").Key("EN_HMM_DIR").SetValue(HMMDirEn)
	cfg.Section("sphinx").Key("EN_DICT_FILE").SetValue(DictFileEn)
	cfg.Section("sphinx").Key("EN_LM_FILE").SetValue(LMFileEn)
	cfg.Section("sphinx").Key("RECORD_THRESHOLD").SetValue(strconv.Itoa(RecordThreshold))
	cfg.Section("sphinx").Key("LOG_FILE").SetValue(SphinxLogFile)

	// Update nlu config
	cfg.Section("nlu").Key("CONVERSATION_API").SetValue(ConversationAPI)
	cfg.Section("nlu").Key("PARSE_API").SetValue(ParseAPI)
	cfg.Section("nlu").Key("PROJECT").SetValue(NluProject)
	cfg.Section("nlu").Key("MODEL").SetValue(NluModel)

	// Update baidu config
	cfg.Section("baidu").Key("ASR_API").SetValue(BaiduASRAPI)
	cfg.Section("baidu").Key("TTS_API").SetValue(BaiduTTSAPI)
	cfg.Section("baidu").Key("VOICE_AUTH_URL").SetValue(BaiduVoiceAuthUrl)
	cfg.Section("baidu").Key("VOICE_API_KEY").SetValue(BaiduVoiceAPIKey)
	cfg.Section("baidu").Key("VOICE_API_SECRET").SetValue(BaiduVoiceAPISecret)

	// Update tts config
	cfg.Section("tts").Key("TTS_DIR").SetValue(TTSDir)
	cfg.Section("tts").Key("TTS_OUT_FILE").SetValue(TTSOutFile)

	if err := cfg.SaveTo(ConfFile); err != nil {
		return fmt.Errorf("Update config file failed: %s", err.Error())
	}
	return nil
}
