//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, August 2019
//

package aiicy

import (
	"context"
	"fmt"
	"github.com/aiicy/aiicy/logger"
	"github.com/aiicy/aiicy/protocol/mqtt"
	"github.com/aiicy/aiicy/sdk/aiicy-go/api"
	"github.com/aiicy/aiicy/utils"
	"io"
	"os"
	"os/signal"
	"syscall"
)

//go:generate mockgen -destination=mock/context.go -package=aiicy github.com/aiicy/aiicy/sdk/aiicy-go Context

// Mode keys
const (
	ModeNative = "native"
	ModeDocker = "docker"
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
	EnvKeyHostID                 = "AIICY_HOST_ID"
	EnvKeyHostOS                 = "AIICY_HOST_OS"
	EnvKeyHostSN                 = "AIICY_HOST_SN"
	EnvKeyWorkDir                = "AIICY_WORK_DIR"
	EnvKeyMasterAPISocket        = "AIICY_MASTER_API_SOCKET"
	EnvKeyMasterGRPCAPISocket    = "AIICY_API_SOCKET"
	EnvKeyMasterAPIAddress       = "AIICY_MASTER_API_ADDRESS"
	EnvKeyMasterGRPCAPIAddress   = "AIICY_API_ADDRESS"
	EnvKeyMasterAPIVersion       = "AIICY_MASTER_API_VERSION"
	EnvKeyServiceMode            = "AIICY_SERVICE_MODE"
	EnvKeyServiceName            = "AIICY_SERVICE_NAME"
	EnvKeyServiceToken           = "AIICY_SERVICE_TOKEN"
	EnvKeyServiceInstanceName    = "AIICY_SERVICE_INSTANCE_NAME"
	EnvKeyServiceInstanceAddress = "AIICY_SERVICE_INSTANCE_ADDRESS"
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
	DefaultBinFile = "bin/aiicy"
	// DefaultBinBackupFile the backup file path of master binary
	DefaultBinBackupFile = "bin/aiicy.old"
	// DefaultSockFile sock file of aiicy by default
	DefaultSockFile = "var/run/aiicy.sock"
	// DefaultGRPCSockFile sock file of grpc api by default
	DefaultGRPCSockFile = "var/run/aiicy/api.sock"
	// DefaultConfFile config path of the service by default
	DefaultConfFile = "etc/aiicy/service.yml"
	// DefaultDBDir db dir of the service by default
	DefaultDBDir = "var/db/aiicy"
	// DefaultRunDir  run dir of the service by default
	DefaultRunDir = "var/run/aiicy"
	// DefaultLogDir  log dir of the service by default
	DefaultLogDir = "var/log/aiicy"
	// DefaultMasterConfDir master config dir by default
	DefaultMasterConfDir = "etc/aiicy"
	// DefaultMasterConfFile master config file by default
	DefaultMasterConfFile = "etc/aiicy/conf.yml"
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
	Log() *logger.Logger
	// check running mode
	IsNative() bool
	// waiting to exit, receiving SIGTERM and SIGINT signals
	Wait()
	// returns wait channel
	WaitChan() <-chan os.Signal

	// Master RESTful API

	// updates application or master
	UpdateSystem(trace, tp, path string) error
	// inspects system stats
	InspectSystem() (*Inspect, error)
	// gets an available port of the host
	GetAvailablePort() (string, error)
	// reports the stats of the instance of the service
	ReportInstance(stats map[string]interface{}) error
	// starts an instance of the service
	StartInstance(serviceName, instanceName string, dynamicConfig map[string]string) error
	// stop the instance of the service
	StopInstance(serviceName, instanceName string) error

	// Master KV API

	// set kv
	SetKV(kv api.KV) error
	// set kv which supports context
	SetKVConext(ctx context.Context, kv api.KV) error
	// get kv
	GetKV(k []byte) (*api.KV, error)
	// get kv which supports context
	GetKVConext(ctx context.Context, k []byte) (*api.KV, error)
	// del kv
	DelKV(k []byte) error
	// del kv which supports context
	DelKVConext(ctx context.Context, k []byte) error
	// list kv with prefix
	ListKV(p []byte) ([]*api.KV, error)
	// list kv with prefix which supports context
	ListKVContext(ctx context.Context, p []byte) ([]*api.KV, error)

	io.Closer
}

type ctx struct {
	sn  string // service name
	in  string // instance name
	md  string // running mode
	cfg ServiceConfig
	log *logger.Logger
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
		log.Errorw(fmt.Sprintf("[%s][%s] failed to create master client", sn, in), logger.Error(err))
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

func (c *ctx) LoadConfig(cfgPath string, cfg interface{}) error {
	if cfgPath == "" {
		cfgPath = DefaultConfFile
	}
	return utils.LoadYAML(cfgPath, cfg)
}

func (c *ctx) Config() *ServiceConfig {
	return &c.cfg
}

func (c *ctx) Log() *logger.Logger {
	return c.log
}

func (c *ctx) Wait() {
	<-c.WaitChan()
	c.Close()
}

func (c *ctx) IsNative() bool {
	return c.md == ModeNative
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
