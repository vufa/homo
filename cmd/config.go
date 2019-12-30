//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, March 2019
//

package cmd

import (
	"fmt"
	"github.com/countstarlight/homo/master"
	"github.com/countstarlight/homo/utils"
	"github.com/go-ini/ini"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

var (
	//App settings
	AppPath    string
	LogPath    string
	OTALogPath string

	// Mode
	DebugMode bool

	// Global config
	Cfg         *ini.File
	ConfFile    string
	WorkDirPath string

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

func LoadConfig() (*master.Config, error) {
	var err error
	if WorkDirPath == "" {
		WorkDirPath, err = WorkDir()
		if err != nil {
			return nil, fmt.Errorf("Get work directory failed: %s", err.Error())
		}
	}

	cfg := &master.Config{}
	// ConfFile = path.Join(WorkDirPath, "conf/app.ini")
	if ConfFile == "" {
		ConfFile = path.Join(WorkDirPath, "conf/conf.yml")
	}

	if !utils.IsFile(ConfFile) {
		return nil, fmt.Errorf("config file %s not found, using default config", ConfFile)
	}
	cfg.File = ConfFile

	err = utils.LoadYAML(cfg.File, cfg)
	if err != nil {
		return nil, fmt.Errorf("Parse config file %s failed: %s", ConfFile, err.Error())
	}

	if err = cfg.Validate(); err != nil {
		return cfg, fmt.Errorf("config invalid: %s", err.Error())
	}

	return cfg, nil
}
