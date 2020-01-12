//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, January 2020
//

package homo

import (
	"encoding/json"
	"github.com/countstarlight/homo/sdk/homo-go/api"
	"os"
)

// Client client of api server
type Client struct {
	*api.Client
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
