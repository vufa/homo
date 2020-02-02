package api

import (
	"time"

	"github.com/aiicy/aiicy/utils"
	"google.golang.org/grpc"
)

const (
	headerKeyUsername = "x-aiicy-username"
	headerKeyPassword = "x-aiicy-password"
)

// ServerConfig api server config
type ServerConfig struct {
	Address           string `yaml:"address" json:"address"`
	utils.Certificate `yaml:",inline" json:",inline"`
}

// ClientConfig api client config
type ClientConfig struct {
	Address           string        `yaml:"address" json:"address"`
	Timeout           time.Duration `yaml:"timeout" json:"timeout" default:"30s"`
	Username          string        `yaml:"username" json:"username"`
	Password          string        `yaml:"password" json:"password"`
	utils.Certificate `yaml:",inline" json:",inline"`
}

// Server server to handle grpc message
type Server struct {
	conf ServerConfig
	svr  *grpc.Server
}

// Client server to handle grpc message
type Client struct {
	conf ClientConfig
	conn *grpc.ClientConn
	KV   KVServiceClient
}
