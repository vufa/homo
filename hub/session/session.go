package session

import (
	"github.com/256dpi/gomqtt/packet"
	"github.com/256dpi/gomqtt/transport"
	"github.com/countstarlight/homo/hub/auth"
	"github.com/countstarlight/homo/hub/common"
	"github.com/countstarlight/homo/logger"
	"github.com/countstarlight/homo/utils"
	"go.uber.org/zap"
	"sync"
)

// session session of a client
// ingress data flow: client -> session(onPublish) -> broker -> database -> session(Ack)
// egress data flow: broker(rule) -> session(doQ0/doQ1) -> client -> session(onPuback)
type session struct {
	id       string
	clean    bool
	clientID string
	conn     transport.Conn
	subs     map[string]packet.Subscription
	manager  *Manager
	pids     *common.PacketIDS
	log      *zap.SugaredLogger
	once     sync.Once
	tomb     utils.Tomb
	sync.Mutex

	authorizer *auth.Authorizer
	//  cache
	permittedPublishTopics map[string]struct{}
}

func newSession(conn transport.Conn, manager *Manager) *session {
	return &session{
		conn:                   conn,
		manager:                manager,
		subs:                   make(map[string]packet.Subscription),
		pids:                   common.NewPacketIDS(),
		log:                    logger.New(logger.LogInfo{Level: "debug"}, "mqtt", "session"),
		permittedPublishTopics: make(map[string]struct{}),
	}
}

// TODO: need to send will message after client reconnected if baetyl panicked
// Situations in which the Will Message is published include, but are not limited to:
// * An I/O error or network failure detected by the Server.
// * The Client fails to communicate within the Keep Alive time.
// * The Client closes the Network Connection without first sending a DISCONNECT Packet. The Server closes the Network Connection because of a protocol error.
func (s *session) sendWillMessage() {
	msg, err := s.manager.recorder.getWill(s.id)
	if err != nil {
		s.log.Errorw("failed to get will message", zap.Error(err))
	}
	if msg == nil {
		return
	}
	err = s.retainMessage(msg)
	if err != nil {
		s.log.Error("failed to retain will message", zap.Error(err))
	}
	s.manager.flow(common.NewMessage(uint32(msg.QOS), msg.Topic, msg.Payload, s.clientID))
}

func (s *session) retainMessage(msg *packet.Message) error {
	if len(msg.Payload) == 0 {
		return s.manager.recorder.removeRetained(msg.Topic)
	}
	return s.manager.recorder.setRetained(msg.Topic, msg)
}

// Close closes this session, only called by session manager
func (s *session) close(will bool) {
	s.once.Do(func() {
		s.tomb.Kill(nil)
		s.log.Infof("session closing, messages (unack): %d", s.pids.Size())
		defer s.log.Infof("session closed, messages (unack): %d", s.pids.Size())
		s.manager.remove(s.id)
		if will {
			s.sendWillMessage()
		}
		s.conn.Close()
		s.manager.recorder.removeWill(s.id)
	})
}
