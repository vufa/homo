//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, December 2019
//

package cmd

import (
	"fmt"
	"github.com/countstarlight/homo/logger"
	"github.com/countstarlight/homo/master"
	"github.com/countstarlight/homo/sdk/homo-go"
	"github.com/countstarlight/homo/utils"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var flags = []cli.Flag{
	&cli.BoolFlag{
		EnvVars:     []string{"HOMO_DEBUG"},
		Name:        "debug",
		Aliases:     []string{"d"},
		Usage:       "start homo in debug mode",
		Destination: &DebugMode,
	},
	&cli.StringFlag{
		EnvVars:     []string{"HOMO_CONFIG_FILE"},
		Name:        "config, c",
		Usage:       "set homo config file path",
		Destination: &ConfFile,
	},
	&cli.StringFlag{
		EnvVars:     []string{"HOMO_WORK_DIR"},
		Name:        "workdir, w",
		Usage:       "set homo work directory",
		Destination: &WorkDirPath,
	},
}

func startInternal(c *cli.Context) error {
	cfg, err := LoadConfig()
	if err != nil {
		return err
	}
	log := logger.New(cfg.Logger, "homo", "master")
	if DebugMode {
		log.Info("in debug mode")
	}
	isOTA := utils.IsFile(cfg.OTALog.Path)
	if isOTA {
		log = logger.New(cfg.OTALog, "type", homo.OTAMST)
	}
	m, err := master.New(WorkDirPath, *cfg, Version, Revision)
	if err != nil {
		log.Errorw("failed to start master", zap.Error(err), zap.String(homo.OTAKeyStep, homo.OTARollingBack))
		/*rberr := master.RollBackMST()
		if rberr != nil {
			log.Errorf("failed to roll back %s", rberr, zap.String(homo.OTAKeyStep, homo.OTAFailure))
			return fmt.Errorf("failed to start master: %s; failed to roll back: %s", err.Error(), rberr.Error())
		}
		log.Infof("master is restarting", zap.String(homo.OTAKeyStep, homo.OTARestarting))*/
		return fmt.Errorf("failed to start master: %s", err.Error())
	}
	return m.Wait()
}
