package session

import (
	"fmt"
	"github.com/256dpi/gomqtt/packet"
	"go.uber.org/zap"
)

// Handle handles mqtt connection
func (s *session) Handle() {
	var err error
	var pkt packet.Generic
	for {
		pkt, err = s.conn.Receive()
		if err != nil {
			if !s.tomb.Alive() {
				return
			}
			s.log.Warnw("failed to reveive message", zap.Error(err))
			s.close(true)
			return
		}
		if _, ok := pkt.(*packet.Connect); !ok && s.authorizer == nil {
			s.log.Errorf("only connect packet is allowed before auth")
			s.close(true)
			return
		}
		switch p := pkt.(type) {
		case *packet.Connect:
			s.log.Debug("received:", p.Type())
			err = s.onConnect(p)
		case *packet.Publish:
			s.log.Debugf("received: %s, pid: %d, qos: %d, topic: %s", p.Type(), p.ID, p.Message.QOS, p.Message.Topic)
			err = s.onPublish(p)
		case *packet.Puback:
			s.log.Debugf("received: %s, pid: %d", p.Type(), p.ID)
			err = s.onPuback(p)
		case *packet.Subscribe:
			s.log.Debugf("received: %s, subs: %v", p.Type(), p.Subscriptions)
			err = s.onSubscribe(p)
		case *packet.Pingreq:
			s.log.Debugln("received:", p.Type())
			err = s.onPingreq(p)
		case *packet.Pingresp:
			s.log.Debugln("received:", p.Type())
			err = nil // just ignore
		case *packet.Disconnect:
			s.log.Debugln("received:", p.Type())
			s.close(false)
			return
		case *packet.Unsubscribe:
			s.log.Debugf("received: %s, topics: %v", p.Type(), p.Topics)
			err = s.onUnsubscribe(p)
		default:
			err = fmt.Errorf("packet (%v) not supported", p)
		}
		if err != nil {
			s.log.Errorf(err.Error())
			s.close(true)
			break
		}
	}
}
