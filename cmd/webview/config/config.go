//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, March 2019
//

package config

import (
	"github.com/countstarlight/homo/module/audio"
	"github.com/countstarlight/homo/module/com"
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

var (
	//App settings
	AppPath string

	// Mode
	DebugMode      bool
	SilenceMode    bool
	OfflineMode    bool
	InterruptMode  bool
	AnalyticalMode bool // display intent and entities

	// Global config
	Cfg      *ini.File
	ConfFile string

	// PortAudio
	RawDir   string
	InputRaw string
	InputWav string
	//sphinx
	HMMDirEn      string
	DictFileEn    string
	LMFileEn      string
	SphinxLogFile string
	// Flag
	IsPlayingVoice bool
	WakeUpd        bool
	WakeUpDisabled bool

	// Lock
	VoicePlayMutex sync.Mutex
	SphinxLoop     sync.WaitGroup // Make sphinx keep capturing audio input
	WakeUpWait     sync.WaitGroup

	RecordThreshold int

	//Auto convert raw pcm buffer to wav
	RawToWav bool
)

const (
	AppName    = "Homo Webview"
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
		logrus.Fatalf("Fail to get app path: %s\n", err.Error())
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

func LoadConfig() {
	workDir, err := WorkDir()
	if err != nil {
		logrus.Fatalf("Fail to get work directory: %s", err.Error())
	}
	ConfFile = path.Join(workDir, "conf/app.ini")

	Cfg, err = ini.Load(ConfFile)
	if err != nil {
		logrus.Fatalf("Fail to parse %s: %s", ConfFile, err.Error())
	}

	Cfg.NameMapper = ini.AllCapsUnderscore

	// Load PortAudio config
	sec := Cfg.Section("portaudio")
	RawDir = sec.Key("RAW_DIR").MustString(path.Join(workDir, "tmp/record"))
	InputRaw = sec.Key("INPUT_RAW").MustString(path.Join(workDir, "tmp/record/input.pcm"))
	InputWav = sec.Key("INPUT_WAV").MustString(path.Join(workDir, "tmp/record/input.wav"))

	// Create path
	if !com.PathExists(RawDir) {
		err := os.MkdirAll(RawDir, os.ModePerm)
		if err != nil {
			logrus.Fatalf("Create path %s failed: %s", RawDir, err.Error())
		}
	}

	// Load sphinx config
	sec = Cfg.Section("sphinx")
	HMMDirEn = sec.Key("EN_HMM_DIR").MustString(path.Join(workDir, "sphinx/en-us/en-us"))
	DictFileEn = sec.Key("EN_DICT_FILE").MustString(path.Join(workDir, "sphinx/homo/homo.dic"))
	LMFileEn = sec.Key("EN_LM_FILE").MustString(path.Join(workDir, "sphinx/homo/homo.lm.bin"))
	SphinxLogFile = sec.Key("LOG_FILE").MustString(path.Join(workDir, "log/sphinx.log"))

	// Update config file
	UpdateConfigFile()
}

func UpdateConfigFile() {
	cfg := ini.Empty()
	if com.IsFile(ConfFile) {
		// Keeps custom settings if there is already something.
		if err := cfg.Append(ConfFile); err != nil {
			logrus.Fatalf("Fail to load conf '%s': %s", ConfFile, err.Error())
		}
	}

	// Update PortAudio config
	cfg.Section("portaudio").Key("RAW_DIR").SetValue(RawDir)
	cfg.Section("portaudio").Key("INPUT_RAW").SetValue(InputRaw)
	cfg.Section("portaudio").Key("INPUT_WAV").SetValue(InputWav)

	// Update sphinx config
	cfg.Section("sphinx").Key("EN_HMM_DIR").SetValue(HMMDirEn)
	cfg.Section("sphinx").Key("EN_DICT_FILE").SetValue(DictFileEn)
	cfg.Section("sphinx").Key("EN_LM_FILE").SetValue(LMFileEn)
	cfg.Section("sphinx").Key("LOG_FILE").SetValue(SphinxLogFile)

	if err := cfg.SaveTo(ConfFile); err != nil {
		logrus.Fatalf("Update config file failed: %s", err.Error())
	}
}

func Terminal(c *cli.Context) error {
	logrus.Infof("退出，开始结束PortAudio...")
	return audio.PaTerminate()
}
