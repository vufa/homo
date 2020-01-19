package api

import (
	"encoding/json"
	"fmt"
	"github.com/countstarlight/homo/master/engine"
	"github.com/countstarlight/homo/protocol/http"
	"github.com/countstarlight/homo/utils"
	"go.uber.org/zap"
	"strconv"
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

// New creates new api server
func New(c http.ServerInfo, m Master, log *zap.SugaredLogger) (*Server, error) {
	svr, err := http.NewServer(c, m.Auth, log)
	if err != nil {
		return nil, err
	}
	s := &Server{
		m: m,
		s: svr,
	}
	s.s.Handle(s.inspectSystem, "GET", "/v1/system/inspect")
	s.s.Handle(s.updateSystem, "PUT", "/v1/system/update")
	s.s.Handle(s.getAvailablePort, "GET", "/v1/ports/available")
	s.s.Handle(s.reportInstance, "PUT", "/v1/services/{serviceName}/instances/{instanceName}/report")
	s.s.Handle(s.startInstance, "PUT", "/v1/services/{serviceName}/instances/{instanceName}/start")
	s.s.Handle(s.stopInstance, "PUT", "/v1/services/{serviceName}/instances/{instanceName}/stop")
	return s, s.s.Start()
}

// Close closes api server
func (s *Server) Close() error {
	return s.s.Close()
}

func (s *Server) inspectSystem(_ http.Params, reqBody []byte) ([]byte, error) {
	return s.m.InspectSystem()
}

func (s *Server) updateSystem(_ http.Params, reqBody []byte) ([]byte, error) {
	if reqBody == nil {
		return nil, fmt.Errorf("request body invalid")
	}
	args := make(map[string]string)
	err := json.Unmarshal(reqBody, &args)
	if err != nil {
		return nil, err
	}
	tp, ok := args["type"]
	if !ok {
		tp = "APP"
	}
	target, ok := args["path"]
	if !ok {
		// backward compatibility, agent version < 0.1.4
		target = args["file"]
	}
	trace, _ := args["trace"]
	go s.m.UpdateSystem(trace, tp, target)
	return nil, nil
}

func (s *Server) getAvailablePort(_ http.Params, reqBody []byte) ([]byte, error) {
	port, err := utils.GetAvailablePort("127.0.0.1")
	if err != nil {
		return nil, err
	}
	res := make(map[string]string)
	res["port"] = strconv.Itoa(port)
	return json.Marshal(res)
}

func (s *Server) reportInstance(params http.Params, reqBody []byte) ([]byte, error) {
	if reqBody == nil {
		return nil, fmt.Errorf("request body invalid")
	}
	serviceName, ok := params["serviceName"]
	if !ok {
		return nil, fmt.Errorf("request params invalid, missing service name")
	}
	instanceName, ok := params["instanceName"]
	if !ok {
		return nil, fmt.Errorf("request params invalid, missing instance name")
	}
	stats := make(map[string]interface{})
	err := json.Unmarshal(reqBody, &stats)
	if err != nil {
		return nil, err
	}
	return nil, s.m.ReportInstance(serviceName, instanceName, stats)
}

func (s *Server) startInstance(params http.Params, reqBody []byte) ([]byte, error) {
	if reqBody == nil {
		return nil, fmt.Errorf("request body invalid")
	}
	serviceName, ok := params["serviceName"]
	if !ok {
		return nil, fmt.Errorf("request params invalid, missing service name")
	}
	instanceName, ok := params["instanceName"]
	if !ok {
		return nil, fmt.Errorf("request params invalid, missing instance name")
	}
	dynamicConfig := make(map[string]string)
	err := json.Unmarshal(reqBody, &dynamicConfig)
	if err != nil {
		return nil, err
	}
	return nil, s.m.StartInstance(serviceName, instanceName, dynamicConfig)
}

func (s *Server) stopInstance(params http.Params, _ []byte) ([]byte, error) {
	serviceName, ok := params["serviceName"]
	if !ok {
		return nil, fmt.Errorf("request params invalid, missing service name")
	}
	instanceName, ok := params["instanceName"]
	if !ok {
		return nil, fmt.Errorf("request params invalid, missing instance name")
	}
	return nil, s.m.StopInstance(serviceName, instanceName)
}
