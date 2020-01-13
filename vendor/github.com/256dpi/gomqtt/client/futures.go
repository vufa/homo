package client

import (
	"time"

	"github.com/256dpi/gomqtt/client/future"
	"github.com/256dpi/gomqtt/packet"
)

// A GenericFuture is returned by publish and unsubscribe methods.
type GenericFuture interface {
	// Wait will wait the given amount of time and return whether the future has
	// been completed, canceled or the request timed out. If no time has been
	// provided the wait will never timeout.

	// Note: Wait will not return any Client related errors.
	Wait(timeout time.Duration) error
}

// A ConnectFuture is returned by the connect method.
type ConnectFuture interface {
	GenericFuture

	// SessionPresent will return whether a session was present.
	SessionPresent() bool

	// ReturnCode will return the connack code returned by the broker.
	ReturnCode() packet.ConnackCode
}

// A SubscribeFuture is returned by the subscribe methods.
type SubscribeFuture interface {
	GenericFuture

	// ReturnCodes will return the suback codes returned by the broker.
	ReturnCodes() []packet.QOS
}

type connectFuture struct {
	*future.Future
}

func (f *connectFuture) SessionPresent() bool {
	// get result
	connack := f.Result().(*packet.Connack)
	if connack == nil {
		return false
	}

	return connack.SessionPresent
}

func (f *connectFuture) ReturnCode() packet.ConnackCode {
	// get result
	connack := f.Result().(*packet.Connack)
	if connack == nil {
		return 0
	}

	return connack.ReturnCode
}

type subscribeFuture struct {
	*future.Future
}

func (f *subscribeFuture) ReturnCodes() []packet.QOS {
	// get result
	suback := f.Result().(*packet.Suback)
	if suback == nil {
		return nil
	}

	return suback.ReturnCodes
}
