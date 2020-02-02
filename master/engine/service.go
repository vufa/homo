package engine

import "github.com/aiicy/aiicy/sdk/aiicy-go"

// Service interfaces of service
type Service interface {
	Name() string
	Engine() Engine
	RestartPolicy() aiicy.RestartPolicyInfo
	Start() error
	Stop()
	Stats()
	StartInstance(instanceName string, dynamicConfig map[string]string) error
	StopInstance(instanceName string) error
}
