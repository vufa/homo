package api

import (
	"github.com/countstarlight/homo/master/engine"
	"github.com/countstarlight/homo/protocol/http"
)

// Master master interface
type Master interface {
	Auth(u, p string) bool

	// for system
	InspectSystem() ([]byte, error)
	UpdateSystem(trace, tp, target string) error

	// for instance
	ReportInstance(serviceName, instanceName string, partialStats engine.PartialStats) error
	StartInstance(serviceName, instanceName string, dynamicConfig map[string]string) error
	StopInstance(serviceName, instanceName string) error
}

// Server master api server
type Server struct {
	m Master
	s *http.Server
}

// Close closes api server
func (s *Server) Close() error {
	return s.s.Close()
}
