package router

import "github.com/countstarlight/homo/hub/common"

// SinkSub subscription of sink
type SinkSub interface {
	ID() string // client id for session
	QOS() uint32
	Topic() string
	TargetQOS() uint32
	TargetTopic() string
	Flow(common.Message)
}

type node struct {
	children map[string]*node
	sinksubs map[string]SinkSub
}

func newNode() *node {
	return &node{
		children: make(map[string]*node),
		sinksubs: make(map[string]SinkSub),
	}
}
