//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that can
// be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, December 2019
//

package cmd

import (
	"fmt"

	"github.com/aiicy/aiicy-go/logger"
	"github.com/aiicy/aiicy/master"
	"github.com/aiicy/aiicy/sdk/aiicy-go"
	"github.com/aiicy/aiicy/utils"
	"github.com/urfave/cli/v2"
)

var flags = []cli.Flag{
	&cli.BoolFlag{
		EnvVars:     []string{"AIICY_DEBUG"},
		Name:        "debug",
		Aliases:     []string{"d"},
		Usage:       "start aiicy in debug mode",
		Destination: &DebugMode,
	},
	&cli.StringFlag{
		EnvVars:     []string{"AIICY_CONFIG_FILE"},
		Name:        "config",
		Aliases:     []string{"c"},
		DefaultText: defaultConfFile,
		Usage:       "set aiicy config file path",
		Destination: &ConfFile,
	},
	&cli.StringFlag{
		EnvVars:     []string{aiicy.EnvKeyWorkDir},
		Name:        "workdir",
		Aliases:     []string{"w"},
		Usage:       "set aiicy work directory",
		Destination: &workDir,
	},
}

func startInternal(c *cli.Context) error {
	cfg, err := LoadConfig()
	if err != nil {
		return err
	}
	var log *logger.Logger
	if DebugMode {
		cfg.Logger.Level = "debug"
		cfg.OTALog.Level = "debug"
		log = logger.New(cfg.Logger, "aiicy", "master")
		log.Info("aiicy running in debug mode")
	} else {
		log = logger.New(cfg.Logger, "aiicy", "master")
	}
	isOTA := utils.IsFile(cfg.OTALog.Path)
	if isOTA {
		log = logger.New(cfg.OTALog, "type", aiicy.OTAMST)
	}
	m, err := master.New(workDir, *cfg, Version, Revision)
	if err != nil {
		log.Errorw("failed to start master", logger.Error(err), logger.String(aiicy.OTAKeyStep, aiicy.OTARollingBack))
		/*rberr := master.RollBackMST()
		if rberr != nil {
			log.Errorf("failed to roll back %s", rberr, logger.String(aiicy.OTAKeyStep, aiicy.OTAFailure))
			return fmt.Errorf("failed to start master: %s; failed to roll back: %s", err.Error(), rberr.Error())
		}
		log.Infof("master is restarting", logger.String(aiicy.OTAKeyStep, aiicy.OTARestarting))*/
		return fmt.Errorf("failed to start master: %s", err.Error())
	}
	return m.Wait()
}
