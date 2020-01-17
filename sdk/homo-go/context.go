//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, August 2019
//

package homo

import (
	"fmt"
	"github.com/countstarlight/homo/logger"
	"github.com/countstarlight/homo/protocol/mqtt"
	"github.com/countstarlight/homo/utils"
	"go.uber.org/zap"
	"io"
	"os"
	"os/signal"
	"syscall"
)

// OTA types
const (
	OTAAPP = "APP"
	OTAMST = "MST"
)

// OTA steps
const (
	OTAKeyStep  = "step"
	OTAKeyType  = "type"
	OTAKeyTrace = "trace"

	OTAReceived    = "RECEIVED"    // [agent] ota event is received
	OTAUpdating    = "UPDATING"    // [master] to update app or master
	OTAUpdated     = "UPDATED"     // [master][finished] app or master is updated
	OTARestarting  = "RESTARTING"  // [master] to restart master
	OTARestarted   = "RESTARTED"   // [master] master is restarted
	OTARollingBack = "ROLLINGBACK" // [master] to roll back app or master
	OTARolledBack  = "ROLLEDBACK"  // [master][finished] app or master is rolled back
	OTAFailure     = "FAILURE"     // [master/agent][finished] failed to update app or master
	OTATimeout     = "TIMEOUT"     // [agent][finished] ota is timed out
)

// CheckOK print OK if binary is valid
const CheckOK = "OK!"

// Env keys
const (
	EnvKeyHostID                 = "HOMO_HOST_ID"
	EnvKeyHostOS                 = "HOMO_HOST_OS"
	EnvKeyHostSN                 = "HOMO_HOST_SN"
	EnvKeyMasterAPISocket        = "HOMO_MASTER_API_SOCKET"
	EnvKeyMasterGRPCAPISocket    = "HOMO_API_SOCKET"
	EnvKeyMasterAPIAddress       = "HOMO_MASTER_API_ADDRESS"
	EnvKeyMasterGRPCAPIAddress   = "HOMO_API_ADDRESS"
	EnvKeyMasterAPIVersion       = "HOMO_MASTER_API_VERSION"
	EnvKeyServiceMode            = "HOMO_SERVICE_MODE"
	EnvKeyServiceName            = "HOMO_SERVICE_NAME"
	EnvKeyServiceToken           = "HOMO_SERVICE_TOKEN"
	EnvKeyServiceInstanceName    = "HOMO_SERVICE_INSTANCE_NAME"
	EnvKeyServiceInstanceAddress = "HOMO_SERVICE_INSTANCE_ADDRESS"
)

// Path keys
const (
	// AppConfFileName application config file name
	AppConfFileName = "application.yml"
	// AppBackupFileName application backup configuration file
	AppBackupFileName = "application.yml.old"
	// AppStatsFileName application stats file name
	AppStatsFileName = "application.stats"
	// MetadataFileName application metadata file name
	MetadataFileName = "metadata.yml"

	// BinFile the file path of master binary
	DefaultBinFile = "bin/homo"
	// DefaultBinBackupFile the backup file path of master binary
	DefaultBinBackupFile = "bin/homo.old"
	// DefaultSockFile sock file of homo by default
	DefaultSockFile = "var/run/homo.sock"
	// DefaultGRPCSockFile sock file of grpc api by default
	DefaultGRPCSockFile = "var/run/homo/api.sock"
	// DefaultConfFile config path of the service by default
	DefaultConfFile = "etc/homo/service.yml"
	// DefaultDBDir db dir of the service by default
	DefaultDBDir = "var/db/homo"
	// DefaultRunDir  run dir of the service by default
	DefaultRunDir = "var/run/homo"
	// DefaultLogDir  log dir of the service by default
	DefaultLogDir = "var/log/homo"
	// DefaultMasterConfDir master config dir by default
	DefaultMasterConfDir = "etc/homo"
	// DefaultMasterConfFile master config file by default
	DefaultMasterConfFile = "etc/homo/conf.yml"

	// backward compatibility
	// PreviousDBDir previous db dir of the service
	PreviousDBDir = "var/db/openedge"
	// PreviousMasterConfDir previous master config dir
	PreviousMasterConfDir = "etc/openedge"
	// PreviousMasterConfFile previous master config file
	PreviousMasterConfFile = "etc/openedge/openedge.yml"
	// PreviousBinBackupFile the backup file path of master binary
	PreviousBinBackupFile = "bin/openedge.old"
	// PreviousLogDir  log dir of the service by default
	PreviousLogDir = "var/log/openedge"
)

// Context of service
type Context interface {
	// returns the system configuration of the service, such as hub and logger
	Config() *ServiceConfig
	// loads the custom configuration of the service
	LoadConfig(string, interface{}) error
	// creates a Client that connects to the Hub through system configuration,
	// you can specify the Client ID and the topic information of the subscription.
	NewHubClient(string, []mqtt.TopicInfo) (*mqtt.Dispatcher, error)
	// returns logger interface
	Log() *zap.SugaredLogger
	// waiting to exit, receiving SIGTERM and SIGINT signals
	Wait()

	// gets an available port of the host
	GetAvailablePort() (string, error)
	// reports the stats of the instance of the service
	ReportInstance(stats map[string]interface{}) error
	// starts an instance of the service
	StartInstance(serviceName, instanceName string, dynamicConfig map[string]string) error
	// stop the instance of the service
	StopInstance(serviceName, instanceName string) error

	io.Closer
}

type ctx struct {
	sn  string // service name
	in  string // instance name
	md  string // running mode
	cfg ServiceConfig
	log *zap.SugaredLogger
	*Client
}

func newContext(s Service) (*ctx, error) {
	var cfg ServiceConfig
	md := os.Getenv(EnvKeyServiceMode)
	sn := os.Getenv(EnvKeyServiceName)
	in := os.Getenv(EnvKeyServiceInstanceName)
	if s.CfgPath == "" {
		s.CfgPath = DefaultConfFile
	}
	err := utils.LoadYAML(s.CfgPath, &cfg)
	if err != nil && !os.IsNotExist(err) {
		logger.S.Fatalf("[%s][%s] failed to load config: %s\n", sn, in, err.Error())
	}
	log := logger.New(cfg.Logger, "service", sn, "instance", in)
	cli, err := NewEnvClient()
	if err != nil {
		log.Errorw(fmt.Sprintf("[%s][%s] failed to create master client", sn, in), zap.Error(err))
	}
	return &ctx{
		sn:     sn,
		in:     in,
		md:     md,
		cfg:    cfg,
		log:    log,
		Client: cli,
	}, nil
}

func (c *ctx) LoadConfig(cfgPath string, cfg interface{}) error {
	if cfgPath == "" {
		cfgPath = DefaultConfFile
	}
	return utils.LoadYAML(cfgPath, cfg)
}

func (c *ctx) NewHubClient(cid string, subs []mqtt.TopicInfo) (*mqtt.Dispatcher, error) {
	if c.cfg.Hub.Address == "" {
		return nil, fmt.Errorf("hub not configured")
	}
	cc := c.cfg.Hub
	if cid != "" {
		cc.ClientID = cid
	}
	if subs != nil {
		cc.Subscriptions = subs
	}
	return mqtt.NewDispatcher(cc, c.log.With("cid", cid)), nil
}

func (c *ctx) Config() *ServiceConfig {
	return &c.cfg
}

func (c *ctx) Log() *zap.SugaredLogger {
	return c.log
}

func (c *ctx) Wait() {
	<-c.WaitChan()
	c.Close()
}

func (c *ctx) WaitChan() <-chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	signal.Ignore(syscall.SIGPIPE)
	return sig
}

func (c *ctx) ReportInstance(stats map[string]interface{}) error {
	return c.Client.ReportInstance(c.sn, c.in, stats)
}

func (c *ctx) Close() error {
	if c.Client.Client != nil {
		if err := c.Client.Close(); err != nil {
			return err
		}
	}
	return nil
}
