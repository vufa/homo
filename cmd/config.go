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
	ConfFile string

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
	AppName = "Homo"
)

// compile variables
var (
	workDir string
	cfgFile string
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

const defaultConfFile = "etc/homo/conf.yml"

func LoadConfig() (*master.Config, error) {
	if ConfFile == "" {
		ConfFile = defaultConfFile
	}

	cfg := &master.Config{File: ConfFile}
	utils.UnmarshalYAML(nil, cfg) // default config
	exe, err := os.Executable()
	if err != nil {
		return cfg, fmt.Errorf("failed to get executable: %s", err.Error())
	}
	exe, err = filepath.EvalSymlinks(exe)
	if err != nil {
		return cfg, fmt.Errorf("failed to get path of executable: %s", err.Error())
	}
	if workDir == "" {
		workDir = path.Dir(path.Dir(exe))
	}
	workDir, err = filepath.Abs(workDir)
	if err != nil {
		return cfg, fmt.Errorf("failed to get absolute path of work directory: %s", err.Error())
	}

	if err = os.Chdir(workDir); err != nil {
		return cfg, fmt.Errorf("failed to change work directory: %s", err.Error())
	}

	if cfgFile != "" {
		cfg.File = cfgFile
	}
	if utils.FileExists(cfg.File) {
		err = utils.LoadYAML(cfg.File, cfg)
		if err != nil {
			return cfg, fmt.Errorf("failed to load config: %s", err.Error())
		}
	} else {
		log.Printf("config file (%s) not found, to use default config", cfg.File)
	}

	if err = cfg.Validate(); err != nil {
		return cfg, fmt.Errorf("config invalid: %s", err.Error())
	}

	return cfg, nil
}
