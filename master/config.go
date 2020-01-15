//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, December 2019
//

package master

import (
	"fmt"
	"github.com/countstarlight/homo/logger"
	"github.com/countstarlight/homo/protocol/http"
	"github.com/countstarlight/homo/sdk/homo-go"
	"github.com/countstarlight/homo/sdk/homo-go/api"
	"github.com/countstarlight/homo/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// DBConf db config
type DBConf struct {
	Driver string
	Path   string
}

// Config master init config
type Config struct {
	Mode     string           `yaml:"mode" json:"mode" default:"docker" validate:"regexp=^(native|docker)$"`
	Server   http.ServerInfo  `yaml:"server" json:"server" default:"{\"address\":\"unix:///var/run/homo.sock\"}"`
	Database DBConf           `yaml:"database" json:"database" default:"{\"driver\":\"sqlite3\",\"path\":\"var/lib/homo/db\"}"`
	API      api.ServerConfig `yaml:"api" json:"api" default:"{\"address\":\"unix:///var/run/homo/api.sock\"}"`
	Logger   logger.LogInfo   `yaml:"logger" json:"logger" default:"{\"path\":\"var/log/homo/homo.log\"}"`
	OTALog   logger.LogInfo   `yaml:"otalog" json:"otalog" default:"{\"path\":\"var/db/homo/ota.log\",\"format\":\"json\"}"`
	Grace    time.Duration    `yaml:"grace" json:"grace" default:"30s"`
	SNFile   string           `yaml:"snfile" json:"snfile"`
	Docker   struct {
		APIVersion string `yaml:"api_version" json:"api_version" default:"1.38"`
	} `yaml:"docker" json:"docker"`
	// cache config file path
	File string
}

// Validate validates config
// TODO: it is not good idea to set envs here
func (c *Config) Validate() error {
	addr := c.Server.Address
	url, err := utils.ParseURL(addr)
	if err != nil {
		return fmt.Errorf("failed to parse address of server: %s", err.Error())
	}
	grpcAddr := c.API.Address
	grpcUrl, err := utils.ParseURL(grpcAddr)
	if err != nil {
		return fmt.Errorf("failed to parse address of API server: %s", err.Error())
	}

	if runtime.GOOS != "linux" && (url.Scheme == "unix" || grpcUrl.Scheme == "unix") {
		return fmt.Errorf("unix domain socket only support on linux, please to use tcp socket")
	}
	if (url.Scheme != "unix" && url.Scheme != "tcp") ||
		(grpcUrl.Scheme != "unix" && grpcUrl.Scheme != "tcp") {
		return fmt.Errorf("only support unix domian socket or tcp socket")
	}

	// address in container
	if url.Scheme == "unix" {
		// http
		sock, err := filepath.Abs(url.Host)
		if err != nil {
			return err
		}
		err = os.MkdirAll(filepath.Dir(sock), 0755)
		if err != nil {
			return err
		}
		utils.SetEnv(homo.EnvKeyMasterAPISocket, sock)
		unixPrefix := "unix://"
		if c.Mode != "native" {
			unixPrefix += "/"
		}
		utils.SetEnv(homo.EnvKeyMasterAPIAddress, unixPrefix+homo.DefaultSockFile)

		// grpc
		grpcSock, err := filepath.Abs(grpcUrl.Host)
		if err != nil {
			return err
		}
		err = os.MkdirAll(filepath.Dir(grpcSock), 0755)
		if err != nil {
			return err
		}
		utils.SetEnv(homo.EnvKeyMasterGRPCAPISocket, grpcSock)
		utils.SetEnv(homo.EnvKeyMasterGRPCAPIAddress, unixPrefix+homo.DefaultGRPCSockFile)
	} else {
		if c.Mode != "native" {
			parts := strings.SplitN(url.Host, ":", 2)
			addr = fmt.Sprintf("tcp://host.docker.internal:%s", parts[1])
			parts = strings.SplitN(grpcUrl.Host, ":", 2)
			grpcUrl.Host = fmt.Sprintf("host.docker.internal:%s", parts[1])
		}
		utils.SetEnv(homo.EnvKeyMasterAPIAddress, addr)
		utils.SetEnv(homo.EnvKeyMasterGRPCAPIAddress, grpcUrl.Host)
	}

	if c.SNFile != "" {
		snByte, err := ioutil.ReadFile(c.SNFile)
		if err != nil {
			fmt.Printf("failed to load SN file: %s", err.Error())
		} else {
			sn := strings.TrimSpace(string(snByte))
			utils.SetEnv(homo.EnvKeyHostSN, sn)
		}
	}

	utils.SetEnv(homo.EnvKeyMasterAPIVersion, "v1")
	utils.SetEnv(homo.EnvKeyHostOS, runtime.GOOS)
	utils.SetEnv(homo.EnvKeyServiceMode, c.Mode)

	hi := utils.GetHostInfo()
	if hi.HostID != "" {
		utils.SetEnv(homo.EnvKeyHostID, hi.HostID)
	}
	return nil
}
