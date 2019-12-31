package engine

import "github.com/countstarlight/homo/sdk/homo-go"

// Service interfaces of service
type Service interface {
	Name() string
	Engine() Engine
	RestartPolicy() homo.RestartPolicyInfo
	Start() error
	Stop()
	Stats()
	StartInstance(instanceName string, dynamicConfig map[string]string) error
	StopInstance(instanceName string) error
}
