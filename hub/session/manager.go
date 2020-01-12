package session

import (
	"github.com/256dpi/gomqtt/transport"
	"github.com/countstarlight/homo/hub/auth"
	"github.com/countstarlight/homo/hub/common"
	"github.com/countstarlight/homo/hub/config"
	"github.com/countstarlight/homo/hub/persist"
	"github.com/countstarlight/homo/hub/rule"
	"github.com/countstarlight/homo/logger"
	cmap "github.com/orcaman/concurrent-map"
	"go.uber.org/zap"
)

// Manager session manager
type Manager struct {
	auth     *auth.Auth
	recorder *recorder
	sessions cmap.ConcurrentMap
	flow     common.Flow
	conf     *config.Message
	rules    *rule.Manager
	log      *zap.SugaredLogger
}

// NewManager creates a session manager
func NewManager(conf *config.Config, flow common.Flow, rules *rule.Manager, pf *persist.Factory) (*Manager, error) {
	sessionDB, err := pf.NewDB("session.db")
	if err != nil {
		return nil, err
	}
	return &Manager{
		auth:     auth.NewAuth(conf.Principals),
		rules:    rules,
		flow:     flow,
		conf:     &conf.Message,
		recorder: newRecorder(sessionDB),
		sessions: cmap.New(),
		log:      logger.New(logger.LogInfo{Level: "debug"}, "manager", "session"),
	}, nil
}

// Handle handles connection
func (m *Manager) Handle(conn transport.Conn) {
	defer conn.Close()
	conn.SetReadLimit(int64(m.conf.Length.Max))
	newSession(conn, m).Handle()
}

// Close closes all sessions, called by rule manager
func (m *Manager) Close() {
	m.log.Infof("session manager closing")
	for item := range m.sessions.IterBuffered() {
		item.Val.(*session).close(true)
	}
	m.log.Infof("session manager closed")
}

// Called by session when error raises
func (m *Manager) remove(id string) {
	m.sessions.Remove(id)
	err := m.rules.RemoveRule(id)
	if err != nil {
		m.log.Debugw("failed to remove rule", zap.Error(err))
	}
}
