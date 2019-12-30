//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, December 2019
//

package cmd

import (
	"github.com/countstarlight/homo/master/config"
	"github.com/countstarlight/homo/sdk/homo-go"
	"github.com/countstarlight/homo/utils/com"
	"github.com/countstarlight/homo/utils/logger"
	"github.com/urfave/cli/v2"
)

var flags = []cli.Flag{
	&cli.BoolFlag{
		EnvVars:     []string{"HOMO_DEBUG"},
		Name:        "debug",
		Aliases:     []string{"d"},
		Usage:       "start homo in debug mode",
		Destination: &config.DebugMode,
	},
	&cli.StringFlag{
		EnvVars:     []string{"HOMO_CONFIG_FILE"},
		Name:        "config, c",
		Usage:       "set homo config file path",
		Destination: &config.ConfFile,
	},
	&cli.StringFlag{
		EnvVars:     []string{"HOMO_WORK_DIR"},
		Name:        "workdir, w",
		Usage:       "set homo work directory",
		Destination: &config.WorkDirPath,
	},
}

func startInternal(c *cli.Context) error {
	if err := config.LoadConfig(); err != nil {
		return err
	}
	log := logger.NewLogger("homo", "master")
	if !config.DebugMode {
		log = logger.NewLoggerToFile(config.LogPath, "homo", "master")
	} else {
		log.Info("in debug mode")
	}
	isOTA := com.IsFile(config.OTALogPath)
	if isOTA {
		log = logger.NewLogger("type", homo.OTAMST)
	}
	return nil
}
