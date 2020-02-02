//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, December 2019
//

package aiicy

import (
	"github.com/aiicy/aiicy/utils"
	"time"
)

// Inspect all aiicy information and status inspected
type Inspect struct {
	// exception information
	Error string `json:"error,omitempty"`
	// inspect time
	Time time.Time `json:"time,omitempty"`
	// software information
	Software Software `json:"software,omitempty"`
	// hardware information
	Hardware Hardware `json:"hardware,omitempty"`
	// service information, including service name, instance running status, etc.
	Services Services `json:"services,omitempty"`
	// storage volume information, including name and version
	Volumes Volumes `json:"volumes,omitempty"`
}

// Software software information
type Software struct {
	// operating system information of host
	OS string `json:"os,omitempty"`
	// CPU information of host
	Arch string `json:"arch,omitempty"`
	// Aiicy process work directory
	PWD string `json:"pwd,omitempty"`
	// Aiicy running mode of application services
	Mode string `json:"mode,omitempty"`
	// Aiicy compiled Golang version
	GoVersion string `json:"go_version,omitempty"`
	// Aiicy release version
	BinVersion string `json:"bin_version,omitempty"`
	// Aiicy git revision
	GitRevision string `json:"git_revision,omitempty"`
	// Aiicy loaded application configuration version
	ConfVersion string `json:"conf_version,omitempty"`
}

// Hardware hardware information
type Hardware struct {
	// host information
	HostInfo *utils.HostInfo `json:"host_stats,omitempty"`
	// net information of host
	NetInfo *utils.NetInfo `json:"net_stats,omitempty"`
	// memory usage information of host
	MemInfo *utils.MemInfo `json:"mem_stats,omitempty"`
	// CPU usage information of host
	CPUInfo *utils.CPUInfo `json:"cpu_stats,omitempty"`
	// disk usage information of host
	DiskInfo *utils.DiskInfo `json:"disk_stats,omitempty"`
	// CPU usage information of host
	GPUInfo *utils.GPUInfo `json:"gpu_stats,omitempty"`
}

// Services all services' information
type Services []ServiceStatus

// ServiceStatus service status
type ServiceStatus struct {
	Name      string           `json:"name,omitempty"`
	Instances []InstanceStatus `json:"instances,omitempty"`
}

// InstanceStatus service instance status
type InstanceStatus map[string]interface{}

// NewInspect create a new information inspected
func NewInspect() *Inspect {
	return &Inspect{
		Services: Services{},
	}
}

// NewServiceStatus create a new service status
func NewServiceStatus(name string) ServiceStatus {
	return ServiceStatus{
		Name:      name,
		Instances: []InstanceStatus{},
	}
}

// Volumes all volumes' information
type Volumes []VolumeStatus

// VolumeStatus volume status
type VolumeStatus struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}
