//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, January 2020
//

package homo

import (
	"encoding/json"
	"fmt"
	"github.com/countstarlight/homo/protocol/http"
	"github.com/countstarlight/homo/sdk/homo-go/api"
	"os"
)

// HTTPClient client of http server
// Deprecated: Use api.Client instead.
type HTTPClient struct {
	cli *http.Client
	ver string
}

// Client client of api server
type Client struct {
	*api.Client
	// TODO: remove
	*HTTPClient
}

// NewEnvClient creates a new client by env
func NewEnvClient() (*Client, error) {
	grpcAddr := os.Getenv(EnvKeyMasterGRPCAPIAddress)
	name := os.Getenv(EnvKeyServiceName)
	token := os.Getenv(EnvKeyServiceToken)

	var (
		gcli *api.Client
		err  error
	)
	if len(grpcAddr) != 0 {
		cc := api.ClientConfig{
			Address:  grpcAddr,
			Username: name,
			Password: token,
		}
		gcli, err = api.NewClient(cc)
		if err != nil {
			return nil, err
		}
	}
	return &Client{
		Client: gcli,
	}, nil
}

// GetAvailablePort gets available port
func (c *Client) GetAvailablePort() (string, error) {
	res, err := c.cli.Get(c.ver + "/ports/available")
	if err != nil {
		return "", err
	}
	info := make(map[string]string)
	err = json.Unmarshal(res, &info)
	if err != nil {
		return "", err
	}
	port, ok := info["port"]
	if !ok {
		return "", fmt.Errorf("invalid response, port not found")
	}
	return port, nil
}

// ReportInstance reports the stats of the instance of the service
func (c *Client) ReportInstance(serviceName, instanceName string, stats map[string]interface{}) error {
	// data, err := json.Marshal(stats)
	_, err := json.Marshal(stats)
	if err != nil {
		return err
	}
	// TODO: gcli reports the stats of the instance of the service
	// _, err = c.cli.Put(data, c.ver+"/services/%s/instances/%s/report", serviceName, instanceName)
	return err
}

// StartInstance starts a new service instance with dynamic config
func (c *Client) StartInstance(serviceName, instanceName string, dynamicConfig map[string]string) error {
	data, err := json.Marshal(dynamicConfig)
	if err != nil {
		return err
	}
	_, err = c.cli.Put(data, c.ver+"/services/%s/instances/%s/start", serviceName, instanceName)
	return err
}

// StopInstance stops a service instance
func (c *Client) StopInstance(serviceName, instanceName string) error {
	_, err := c.cli.Put(nil, c.ver+"/services/%s/instances/%s/stop", serviceName, instanceName)
	return err
}
