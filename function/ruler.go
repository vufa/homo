package main

import (
	"github.com/256dpi/gomqtt/packet"
	"github.com/countstarlight/homo/protocol/mqtt"
	"github.com/countstarlight/homo/sdk/homo-go"
	"github.com/docker/distribution/uuid"
	"go.uber.org/zap"
	"time"
)

type ruler struct {
	cfg RuleInfo
	fun *Function
	hub *mqtt.Dispatcher
	log *zap.SugaredLogger
}

func newRuler(ri RuleInfo, c *mqtt.Dispatcher, f *Function, log *zap.SugaredLogger) *ruler {
	return &ruler{
		cfg: ri,
		hub: c,
		fun: f,
		log: log.With("rule", ri.ClientID),
	}
}

func (rr *ruler) start() error {
	return rr.hub.Start(rr)
}

func (rr *ruler) ProcessPublish(pkt *packet.Publish) error {
	msg := &homo.FunctionMessage{
		ID:               uint64(pkt.ID),
		QOS:              uint32(pkt.Message.QOS),
		Topic:            pkt.Message.Topic,
		Payload:          pkt.Message.Payload,
		Timestamp:        time.Now().UTC().Unix(),
		FunctionName:     rr.cfg.Function.Name,
		FunctionInvokeID: uuid.Generate().String(),
	}
	return rr.fun.CallAsync(msg, rr.callback)
}

func (rr *ruler) ProcessPuback(pkt *packet.Puback) error {
	return rr.hub.Send(pkt)
}

func (rr *ruler) ProcessError(err error) {
	rr.log.Errorf(err.Error())
}

func (rr *ruler) close() {
	rr.hub.Close()
}
