//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, August 2019
//

package homo

import (
	"time"
)

// RestartPolicyInfo holds the policy of a module
type RestartPolicyInfo struct {
	Retry struct {
		Max int `yaml:"max" json:"max"`
	} `yaml:"retry" json:"retry"`
	Policy  string      `yaml:"policy" json:"policy" default:"always"`
	Backoff BackoffInfo `yaml:"backoff" json:"backoff"`
}

// BackoffInfo holds backoff value
type BackoffInfo struct {
	Min    time.Duration `yaml:"min" json:"min" default:"1s" validate:"min=1000000000"`
	Max    time.Duration `yaml:"max" json:"max" default:"5m" validate:"min=1000000000"`
	Factor float64       `yaml:"factor" json:"factor" default:"2" validate:"min=1"`
}

// Resources resources config
type Resources struct {
	CPU    CPU    `yaml:"cpu" json:"cpu"`
	Pids   Pids   `yaml:"pids" json:"pids"`
	Memory Memory `yaml:"memory" json:"memory"`
}

// CPU cpu config
type CPU struct {
	Cpus    float64 `yaml:"cpus" json:"cpus"`
	SetCPUs string  `yaml:"setcpus" json:"setcpus"`
}

// Pids pids config
type Pids struct {
	Limit int64 `yaml:"limit" json:"limit"`
}

// Memory memory config
type Memory struct {
	Limit int64 `yaml:"limit" json:"limit"`
	Swap  int64 `yaml:"swap" json:"swap"`
}

type memory struct {
	Limit string `yaml:"limit" json:"limit"`
	Swap  string `yaml:"swap" json:"swap"`
}

// NetworksInfo network configurations of service
type NetworksInfo struct {
	ServiceNetworks map[string]ServiceNetwork `yaml:"networks" json:"networks"`
}

// ServiceNetwork specific network configuration of service
type ServiceNetwork struct {
	Aliases     []string `yaml:"aliases" json:"aliases"`
	Ipv4Address string   `yaml:"ipv4_address" json:"ipv4_address"`
}

// ComposeAppConfig application configuration of compose
type ComposeAppConfig struct {
	// specifies the version of compose file
	Version string `yaml:"version,omitempty" json:"version,omitempty"`
	// specifies name of the application
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
	// specifies the app version of the application configuration
	AppVersion string `yaml:"app_version,omitempty" json:"app_version,omitempty"`
	// specifies the service information of the application
	Services map[string]ComposeService `yaml:"services,omitempty" json:"services,omitempty" default:"{}"`
	// specifies the storage volume information of the application
	Volumes map[string]ComposeVolume `yaml:"volumes,omitempty" json:"volumes,omitempty" default:"{}"`
	// specifies the network information of the application
	Networks map[string]ComposeNetwork `yaml:"networks,omitempty" json:"networks,omitempty" default:"{}"`
}

// ComposeService service configuration of compose
type ComposeService struct {
	// specifies the unique name of the service
	ContainerName string `yaml:"container_name,omitempty" json:"container_name,omitempty"`
	// specifies the hostname of the service
	Hostname string `yaml:"hostname,omitempty" json:"hostname,omitempty"`
	// specifies the image of the service, usually using the Docker image name
	Image string `yaml:"image,omitempty" json:"image,omitempty" validate:"nonzero"`
	// specifies the number of instances started
	Replica int `yaml:"replica,omitempty" json:"replica,omitempty" validate:"min=0"`
	// specifies the storage volumes that the service needs, map the storage volume to the directory in the container
	Volumes []*ServiceVolume `yaml:"volumes,omitempty" json:"volumes,omitempty"`
	// specifies the network mode of the service
	NetworkMode string `yaml:"network_mode,omitempty" json:"network_mode,omitempty" validate:"regexp=^(bridge|host|none)?$"`
	// specifies the network that the service needs
	Networks NetworksInfo `yaml:"networks,omitempty" json:"networks,omitempty"`
	// specifies the port bindings which exposed by the service, only for Docker container mode
	Ports []string `yaml:"ports,omitempty" json:"ports,omitempty" default:"[]"`
	// specifies the device bindings which used by the service, only for Docker container mode
	Devices []string `yaml:"devices,omitempty" json:"devices,omitempty" default:"[]"`
	// specified other depended services
	DependsOn []string `yaml:"depends_on,omitempty" json:"depends_on,omitempty" default:"[]"`
	// specifies the startup arguments of the service program, but does not include `arg[0]`
	Command *Command `yaml:"command,omitempty" json:"command,omitempty"`
	// specifies the environment variable of the service program
	Environment *Environment `yaml:"environment,omitempty" json:"environment,omitempty" default:"{}"`
	// specifies the restart policy of the instance of the service
	Restart RestartPolicyInfo `yaml:"restart,omitempty" json:"restart,omitempty"`
	// specifies resource limits for a single instance of the service,  only for Docker container mode
	Resources Resources `yaml:"resources,omitempty" json:"resources,omitempty"`
	// specifies runtime to use, only for Docker container mode
	Runtime string `yaml:"runtime,omitempty" json:"runtime,omitempty"`
}

// ComposeVolume volume configuration of compose
type ComposeVolume struct {
	// specified driver for the storage volume
	Driver string `yaml:"driver,omitempty" json:"driver,omitempty" default:"local"`
	// specified driver options for the storage volume
	DriverOpts map[string]string `yaml:"driver_opts,omitempty" json:"driver_opts,omitempty" default:"{}"`
	// specified labels for the storage volume
	Labels map[string]string `yaml:"labels,omitempty" json:"labels,omitempty" default:"{}"`
}

// ComposeNetwork network configuration
type ComposeNetwork struct {
	// specifies driver for network
	Driver string `yaml:"driver,omitempty" json:"driver,omitempty" default:"bridge"`
	// specified driver options for network
	DriverOpts map[string]string `yaml:"driver_opts,omitempty" json:"driver_opts,omitempty" default:"{}"`
	// specifies labels to add metadata
	Labels map[string]string `yaml:"labels,omitempty" json:"labels,omitempty" default:"{}"`
}

// Environment environment
type Environment struct {
	Envs map[string]string `yaml:"envs" json:"envs" default:"{}"`
}

// ServiceVolume specific volume configuration of service
type ServiceVolume struct {
	// specifies type of volume
	Type string `yaml:"type,omitempty" json:"type,omitempty" validate:"regexp=^(volume|bind)?$"`
	// specifies source of volume
	Source string `yaml:"source,omitempty" json:"source,omitempty"`
	// specifies target of volume
	Target string `yaml:"target,omitempty" json:"target,omitempty"`
	// specifies if the volume is read-only
	ReadOnly bool `yaml:"read_only,omitempty" json:"read_only,omitempty"`
}

// Command command configuration of the service
type Command struct {
	Cmd []string `yaml:"cmd" json:"cmd" default:"[]"`
}
