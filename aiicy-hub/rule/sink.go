package rule

import (
	"sync/atomic"
	"time"

	"github.com/aiicy/aiicy-go/logger"
	"github.com/aiicy/aiicy/aiicy-hub/common"
	"github.com/aiicy/aiicy/aiicy-hub/router"
	"github.com/aiicy/aiicy/utils"
)

type sink struct {
	id      string
	offset  uint64
	broker  broker
	msgchan *msgchan
	trieq0  *router.Trie
	trieq1  *router.Trie
	tomb    utils.Tomb
	log     *logger.Logger
}

func newSink(id string, b broker, r *router.Trie, msgchan *msgchan, log *logger.Logger) *sink {
	s := &sink{
		id:      id,
		broker:  b,
		trieq0:  r,
		trieq1:  router.NewTrie(),
		msgchan: msgchan,
		log:     log.With("sink", id),
	}
	return s
}

func (s *sink) getOffset() uint64 {
	return atomic.LoadUint64(&s.offset)
}

func (s *sink) setOffset(v uint64) {
	atomic.StoreUint64(&s.offset, v)
}

// Register adds a subscription
func (s *sink) register(sub *sinksub) {
	s.trieq0.Add(sub)
	s.trieq1.Add(sub)
}

// Remove removes a subscription
func (s *sink) remove(id, topic string) {
	s.trieq0.Remove(id, topic)
	s.trieq1.Remove(id, topic)
}

// RemoveAll removes all subscriptions by id
func (s *sink) removeAll(id string) {
	s.trieq0.RemoveAll(id)
}

func (s *sink) start() error {
	if s.id == common.RuleMsgQ0 {
		return s.tomb.Go(s.goRoutingQ0)
	}

	offset, err := s.broker.InitOffset(s.id, s.msgchan.persist != nil)
	if err != nil {
		return err
	}
	s.setOffset(offset)
	return s.tomb.Go(s.goRoutingQ1)
}

func (s *sink) stop() {
	s.log.Debugf("sink stopping")
	s.trieq0.RemoveAll(s.id)
	s.tomb.Kill(nil)
}

func (s *sink) wait() {
	err := s.tomb.Wait()
	s.log.Debugw("sink stopped", logger.Error(err))
}

func (s *sink) goRoutingQ0() error {
	s.log.Debug("task of routing message (Q0) begins")
	defer s.log.Debug("task of routing message (Q0) stopped")

	var msg *common.Message
	for {
		select {
		case <-s.tomb.Dying():
			return nil
		case msg = <-s.broker.MsgQ0Chan():
			matches := s.trieq0.MatchUnique(msg.Topic)
			for _, sub := range matches {
				sub.Flow(*msg)
			}
		}
	}
}

func (s *sink) goRoutingQ1() error {
	s.log.Debugf("task of routing message (Q1) begins with offset=%d", s.getOffset())
	defer s.log.Debug("task of routing message (Q1) stopped")

	var (
		err    error
		msg    *common.Message
		msgs   []*common.Message
		length int
	)
	ticker := time.NewTicker(time.Millisecond * 10)
	maxBatchSize := s.broker.Config().Message.Egress.Qos1.Batch.Max
	for {
		if !s.tomb.Alive() {
			return nil
		}
		msgs, err = s.broker.FetchQ1(s.getOffset(), maxBatchSize)
		if err != nil {
			s.log.Errorw("failed to fetch message", logger.Error(err))
			select {
			case <-s.tomb.Dying():
				return nil
			case <-time.After(time.Second):
				continue
			}
		}
		length = len(msgs)
		if length == 0 {
			select {
			case <-s.tomb.Dying():
				return nil
			case <-ticker.C:
				continue
			}
		}
		s.log.Debugf("%d message(s) fetched", length)
		if length != 1 {
			for _, msg = range msgs[:length-1] {
				matches := s.trieq1.MatchUnique(msg.Topic)
				for _, sub := range matches {
					sub.Flow(*msg)
				}
			}
		}
		msg = msgs[length-1]
		matches := s.trieq1.MatchUnique(msg.Topic)
		for _, sub := range matches {
			sub.Flow(*msg)
		}
		if len(matches) == 0 {
			// put barrier to make sure offset update in db even no message routed
			msg.Barrier = true
			s.msgchan.putQ0(msg)
		}
		s.setOffset(msg.SequenceID + 1)
	}
}
