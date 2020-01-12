package rule

import (
	"github.com/countstarlight/homo/hub/router"
	"github.com/countstarlight/homo/logger"
	"github.com/countstarlight/homo/utils"
	"go.uber.org/zap"
)

type sink struct {
	id      string
	offset  uint64
	broker  broker
	msgchan *msgchan
	trieq0  *router.Trie
	trieq1  *router.Trie
	tomb    utils.Tomb
	log     *zap.SugaredLogger
}

func newSink(id string, b broker, r *router.Trie, msgchan *msgchan) *sink {
	s := &sink{
		id:      id,
		broker:  b,
		trieq0:  r,
		trieq1:  router.NewTrie(),
		msgchan: msgchan,
		log:     logger.New(logger.LogInfo{Level: "debug"}, "sink", id),
	}
	return s
}
