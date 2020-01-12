package rule

import (
	"github.com/countstarlight/homo/hub/common"
	"github.com/countstarlight/homo/utils"
	"github.com/jpillora/backoff"
	"go.uber.org/zap"
	"time"
)

// Q0: it means the message with QOS=0 published by clients or functions
// Q1: it means the message with QOS=1 published by clients or functions

// msgchan message channel routed from sink
type msgchan struct {
	msgq0            chan *common.Message
	msgq1            chan *common.Message
	msgack           chan *common.MsgAck // TODO: move to sink?
	persist          func(uint64)
	msgtomb          utils.Tomb
	acktomb          utils.Tomb
	quitTimeout      time.Duration
	publish          common.Publish
	republish        common.Publish
	republishBackoff *backoff.Backoff
	log              *zap.SugaredLogger
}

// newMsgChan creates a new message channel
func newMsgChan(l0, l1 int, publish, republish common.Publish, republishTimeout time.Duration, quitTimeout time.Duration, persist func(uint64), log *zap.SugaredLogger) *msgchan {
	backoff := &backoff.Backoff{
		Min:    time.Millisecond * 100,
		Max:    republishTimeout,
		Factor: 2,
	}
	return &msgchan{
		msgq0:            make(chan *common.Message, l0),
		msgq1:            make(chan *common.Message, l1),
		msgack:           make(chan *common.MsgAck, l1),
		publish:          publish,
		republish:        republish,
		republishBackoff: backoff,
		quitTimeout:      quitTimeout,
		persist:          persist,
		log:              log,
	}
}
