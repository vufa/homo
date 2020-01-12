package rule

import (
	"github.com/countstarlight/homo/hub/common"
	"github.com/countstarlight/homo/hub/router"
	"github.com/countstarlight/homo/logger"
	"go.uber.org/zap"
	"sync"
)

type base interface {
	uid() string
	start() (err error)
	stop()
	wait(bool)
	channel() *msgchan
	register(sub *sinksub)
	remove(id, topic string)
	info() map[string]interface{}
}

type rulebase struct {
	id      string
	sink    *sink
	broker  broker
	msgchan *msgchan
	once    sync.Once
	log     *zap.SugaredLogger
}

func newRuleBase(id string, persistent bool, b broker, r *router.Trie, publish, republish common.Publish) *rulebase {
	log := logger.New(logger.LogInfo{Level: "debug"}, "rule", id)
	rb := &rulebase{
		id:     id,
		broker: b,
		log:    log,
	}
	persist := rb.persist
	if !persistent {
		persist = nil
	}
	rb.msgchan = newMsgChan(
		b.Config().Message.Egress.Qos0.Buffer.Size,
		b.Config().Message.Egress.Qos1.Buffer.Size,
		publish,
		republish,
		b.Config().Message.Egress.Qos1.Retry.Interval,
		b.Config().Shutdown.Timeout,
		persist,
		log,
	)
	rb.sink = newSink(id, b, r, rb.msgchan)
	return rb
}

func newRuleQos0(b broker, r *router.Trie) *rulebase {
	return newRuleBase(common.RuleMsgQ0, false, b, r, nil, nil)
}

func newRuleTopic(b broker, r *router.Trie) *rulebase {
	rb := newRuleBase(common.RuleTopic, true, b, r, nil, nil)
	rb.msgchan.publish = rb.publish
	return rb
}

func (r *rulebase) publish(msg common.Message) {
	msg.QOS = msg.TargetQOS
	msg.Topic = msg.TargetTopic
	msg.SequenceID = 0
	if msg.QOS == 1 {
		msg.SetCallbackPID(0, func(_ uint32) { msg.Ack() })
	}
	r.broker.Flow(&msg)
}

func (r *rulebase) persist(sid uint64) {
	err := r.broker.PersistOffset(r.id, sid)
	if err != nil {
		r.log.Errorw("failed to persist offset", zap.Error(err))
	}
}
