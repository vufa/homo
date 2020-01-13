package mqtt

import "github.com/256dpi/gomqtt/packet"

// Handler MQTT message handler interface
type Handler interface {
	ProcessPublish(*packet.Publish) error
	ProcessPuback(*packet.Puback) error
	ProcessError(error)
}
