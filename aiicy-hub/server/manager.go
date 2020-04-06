package server

import (
	"github.com/256dpi/gomqtt/transport"
	"github.com/aiicy/aiicy-go/logger"
	"github.com/aiicy/aiicy/protocol/mqtt"
	"github.com/aiicy/aiicy/utils"
)

// Handle handles connection
type Handle func(transport.Conn)

// Manager manager of servers
type Manager struct {
	servers []transport.Server
	handle  Handle
	tomb    utils.Tomb
	log     *logger.Logger
}

// NewManager creates a server manager
func NewManager(addrs []string, cert utils.Certificate, handle Handle, log *logger.Logger) (*Manager, error) {
	launcher, err := mqtt.NewLauncher(cert)
	if err != nil {
		return nil, err
	}
	m := &Manager{
		servers: make([]transport.Server, 0),
		handle:  handle,
		log:     log.With("manager", "server"),
	}
	for _, addr := range addrs {
		svr, err := launcher.Launch(addr)
		if err != nil {
			m.Close()
			return nil, err
		}
		m.servers = append(m.servers, svr)
	}
	return m, nil
}

// Start starts all servers
func (m *Manager) Start() {
	for _, item := range m.servers {
		svr := item
		m.tomb.Go(func() error {
			for {
				conn, err := svr.Accept()
				if err != nil {
					if !m.tomb.Alive() {
						return nil
					}
					m.log.Errorw("failed to accept connection", logger.Error(err))
					continue
				}
				go m.handle(conn)
			}
		})
	}
}

// Close closes server manager
func (m *Manager) Close() error {
	m.tomb.Kill(nil)
	for _, svr := range m.servers {
		svr.Close()
	}
	return m.tomb.Wait()
}
