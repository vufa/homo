package session

import (
	"encoding/json"
	"fmt"
	"github.com/256dpi/gomqtt/packet"
	"github.com/countstarlight/homo/hub/common"
	"github.com/countstarlight/homo/hub/persist"
	"github.com/countstarlight/homo/logger"
	"go.uber.org/zap"
	"sync"
)

// recorder records session info
type recorder struct {
	db  persist.Database
	log *zap.SugaredLogger
	sync.Mutex
}

// NewRecorder creates a recorder
func newRecorder(db persist.Database) *recorder {
	return &recorder{
		db:  db,
		log: logger.New(logger.LogInfo{Level: "debug"}, "session", "recorder"),
	}
}

// SetRetained sets retained message of topic
func (c *recorder) setRetained(topic string, msg *packet.Message) error {
	c.Lock()
	defer c.Unlock()
	if !msg.Retain {
		return nil
	}
	value, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	err = c.db.BucketPut(common.BucketNameDotRetained, []byte(topic), value)
	if err != nil {
		c.log.Errorw(fmt.Sprintf("failed to persist retain message: topic=%s", topic), zap.Error(err))
		return fmt.Errorf("failed to persist retain message: %s", err.Error())
	}
	c.log.Debugf("retain message persisited: topic=%s", topic)
	return nil
}

// RemoveRetained removes retain message of topic
func (c *recorder) removeRetained(topic string) error {
	err := c.db.BucketDelete(common.BucketNameDotRetained, []byte(topic))
	if err != nil {
		c.log.Errorw(fmt.Sprintf("failed to remove retain message: topic=%s", topic), zap.Error(err))
		return fmt.Errorf("failed to remove retain message: %s", err.Error())
	}
	c.log.Debugf("retain message removed: topic=%s", topic)
	return nil
}

// GetWill gets will message of cleint
func (c *recorder) getWill(id string) (*packet.Message, error) {
	data, err := c.db.BucketGet(common.BucketNameDotWill, []byte(id))
	if err != nil {
		c.log.Errorw(fmt.Sprintf("failed to get will message: id=%s", id), zap.Error(err))
		return nil, fmt.Errorf("failed to get will message: %s", err.Error())
	}
	if len(data) == 0 {
		return nil, nil
	}
	var msg packet.Message
	err = json.Unmarshal(data, &msg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal will message: %s", err.Error())
	}
	c.log.Debugf("will message got: topic=%s, id=%s", msg.Topic, id)
	return &msg, nil
}

// RemoveWill removes will message of cleint
func (c *recorder) removeWill(id string) error {
	err := c.db.BucketDelete(common.BucketNameDotWill, []byte(id))
	if err != nil {
		c.log.Errorw(fmt.Sprintf("failed to remove will message: id=%s", id), zap.Error(err))
		return fmt.Errorf("failed to remove will message: %s", err.Error())
	}
	c.log.Debugf("will message removed: id=%s", id)
	return nil
}
