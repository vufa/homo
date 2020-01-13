//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, December 2019
//

package main

import (
	"github.com/countstarlight/homo/hub/broker"
	"github.com/countstarlight/homo/hub/config"
	"github.com/countstarlight/homo/hub/persist"
	"github.com/countstarlight/homo/hub/rule"
	"github.com/countstarlight/homo/hub/server"
	"github.com/countstarlight/homo/hub/session"
	"github.com/countstarlight/homo/logger"
	"github.com/countstarlight/homo/sdk/homo-go"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"os"
)

type mo struct {
	ctx      homo.Context
	cfg      config.Config
	cfgPath  string
	Rules    *rule.Manager
	Sessions *session.Manager
	broker   *broker.Broker
	servers  *server.Manager
	factory  *persist.Factory
	log      *zap.SugaredLogger
}

func (m *mo) start() error {
	err := m.ctx.LoadConfig(m.cfgPath, &m.cfg)
	if err != nil {
		m.log.Errorw("failed to load config:", zap.Error(err))
		return err
	}
	m.factory, err = persist.NewFactory(m.cfg.Storage.Dir)
	if err != nil {
		m.log.Errorw("failed to new factory:", zap.Error(err))
		return err
	}
	m.broker, err = broker.NewBroker(&m.cfg, m.factory, m.ctx.ReportInstance, m.log)
	if err != nil {
		m.log.Errorw("failed to new broker:", zap.Error(err))
		return err
	}
	m.Rules, err = rule.NewManager(m.cfg.Subscriptions, m.broker, m.ctx.ReportInstance, m.log)
	if err != nil {
		m.log.Errorw("failed to new rule manager:", zap.Error(err))
		return err
	}
	m.Sessions, err = session.NewManager(&m.cfg, m.broker.Flow, m.Rules, m.factory, m.log)
	if err != nil {
		m.log.Errorw("failed to new session manager:", zap.Error(err))
		return err
	}
	m.servers, err = server.NewManager(m.cfg.Listen, m.cfg.Certificate, m.Sessions.Handle, m.log)
	if err != nil {
		m.log.Errorw("failed to new server manager:", zap.Error(err))
		return err
	}
	m.Rules.Start()
	m.servers.Start()
	return nil
}

func (m *mo) close() {
	if m.Rules != nil {
		m.Rules.Close()
	}
	if m.Sessions != nil {
		m.Sessions.Close()
	}
	if m.servers != nil {
		m.servers.Close()
	}
	if m.broker != nil {
		m.broker.Close()
	}
	if m.factory != nil {
		m.factory.Close()
	}
}

func main() {
	var cfgPath string
	hub := &cli.App{
		Name:    "Homo Hub",
		Version: "0.0.1",
		Usage:   "Hub for homo",
		Flags: []cli.Flag{
			&cli.StringFlag{
				EnvVars:     []string{"HOMO_HUB_CONFIG_FILE"},
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "set homo hub config file path",
				Destination: &cfgPath,
			},
		},
		Action: func(c *cli.Context) error {
			homo.Run(homo.Service{CfgPath: cfgPath}, func(ctx homo.Context) error {
				m := mo{ctx: ctx, log: ctx.Log(), cfgPath: cfgPath}
				defer m.close()
				err := m.start()
				if err != nil {
					return err
				}
				ctx.Wait()
				return nil
			})
			return nil
		},
	}
	if err := hub.Run(os.Args); err != nil {
		logger.S.Fatal(err)
	}
}
