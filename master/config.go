//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that can
// be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, December 2019
//

package master

import (
	"fmt"
	"github.com/aiicy/aiicy/sdk/aiicy-go"
	"github.com/aiicy/aiicy/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

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
		utils.SetEnv(aiicy.EnvKeyMasterAPISocket, sock)
		unixPrefix := "unix://"
		if c.Mode != "native" {
			unixPrefix += "/"
		}
		utils.SetEnv(aiicy.EnvKeyMasterAPIAddress, unixPrefix+aiicy.DefaultSockFile)

		// grpc
		grpcSock, err := filepath.Abs(grpcUrl.Host)
		if err != nil {
			return err
		}
		err = os.MkdirAll(filepath.Dir(grpcSock), 0755)
		if err != nil {
			return err
		}
		utils.SetEnv(aiicy.EnvKeyMasterGRPCAPISocket, grpcSock)
		utils.SetEnv(aiicy.EnvKeyMasterGRPCAPIAddress, unixPrefix+aiicy.DefaultGRPCSockFile)
	} else {
		if c.Mode != "native" {
			parts := strings.SplitN(url.Host, ":", 2)
			addr = fmt.Sprintf("tcp://host.docker.internal:%s", parts[1])
			parts = strings.SplitN(grpcUrl.Host, ":", 2)
			grpcUrl.Host = fmt.Sprintf("host.docker.internal:%s", parts[1])
		}
		utils.SetEnv(aiicy.EnvKeyMasterAPIAddress, addr)
		utils.SetEnv(aiicy.EnvKeyMasterGRPCAPIAddress, grpcUrl.Host)
	}

	if c.SNFile != "" {
		snByte, err := ioutil.ReadFile(c.SNFile)
		if err != nil {
			fmt.Printf("failed to load SN file: %s", err.Error())
		} else {
			sn := strings.TrimSpace(string(snByte))
			utils.SetEnv(aiicy.EnvKeyHostSN, sn)
		}
	}

	utils.SetEnv(aiicy.EnvKeyMasterAPIVersion, "v1")
	utils.SetEnv(aiicy.EnvKeyHostOS, runtime.GOOS)
	utils.SetEnv(aiicy.EnvKeyServiceMode, c.Mode)

	hi := utils.GetHostInfo()
	if hi.HostID != "" {
		utils.SetEnv(aiicy.EnvKeyHostID, hi.HostID)
	}
	return nil
}
