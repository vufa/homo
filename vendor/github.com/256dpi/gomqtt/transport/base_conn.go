package transport

import (
	"io"
	"sync"
	"time"

	"github.com/256dpi/gomqtt/packet"
)

// A Carrier is a generalized stream that can be used with BaseConn.
type Carrier interface {
	io.ReadWriteCloser
	SetReadDeadline(time.Time) error
}

// A BaseConn manages the low-level plumbing between the Carrier and the packet
// Stream.
type BaseConn struct {
	carrier      Carrier
	stream       *packet.Stream
	sendMutex    sync.Mutex
	receiveMutex sync.Mutex
	readTimeout  time.Duration
}

// NewBaseConn creates a new BaseConn using the specified Carrier.
func NewBaseConn(c Carrier) *BaseConn {
	return &BaseConn{
		carrier: c,
		stream:  packet.NewStream(c, c),
	}
}

// Send will write the packet to an internal buffer. It will either flush the
// internal buffer immediately or asynchronously in the background when it gets
// stale. Encoding errors are directly returned, but any network errors caught
// while flushing the buffer asynchronously will be returned on the next call.
//
// Note: Only one goroutine can send at the same time.
func (c *BaseConn) Send(pkt packet.Generic, async bool) error {
	// acquire mutex
	c.sendMutex.Lock()
	defer c.sendMutex.Unlock()

	// write packet
	err := c.stream.Write(pkt, async)
	if err != nil {
		// ensure carrier gets closed
		_ = c.carrier.Close()

		return err
	}

	return nil
}

// Receive will read from the underlying connection and return a fully read
// packet. It will return any error encountered while decoding or reading from
// the underlying connection.
//
// Note: Only one goroutine can receive at the same time.
func (c *BaseConn) Receive() (packet.Generic, error) {
	// acquire mutex
	c.receiveMutex.Lock()
	defer c.receiveMutex.Unlock()

	// read next packet
	pkt, err := c.stream.Read()
	if err != nil {
		// ensure carrier gets closed
		_ = c.carrier.Close()

		return nil, err
	}

	// reset timeout
	err = c.resetTimeout()
	if err != nil {
		// ensure carrier is closed
		_ = c.carrier.Close()

		return nil, err
	}

	return pkt, nil
}

// Close will close the underlying connection and cleanup resources. It will
// return any error encountered while closing the underlying connection.
func (c *BaseConn) Close() error {
	// acquire mutex
	c.sendMutex.Lock()
	defer c.sendMutex.Unlock()

	// flush buffer
	err1 := c.stream.Flush()

	// close carrier
	err2 := c.carrier.Close()

	// handle errors
	if err1 != nil {
		return err1
	} else if err2 != nil {
		return err2
	}

	return nil
}

// SetReadLimit sets the maximum size of a packet that can be received.
// If the limit is greater than zero, Receive will close the connection and
// return an error if receiving the next packet will exceed the limit.
func (c *BaseConn) SetReadLimit(limit int64) {
	c.stream.SetReadLimit(limit)
}

// SetReadTimeout sets the maximum time that can pass between reads.
// If no data is received in the set duration the connection will be closed
// and Read returns an error.
func (c *BaseConn) SetReadTimeout(timeout time.Duration) {
	// acquire mutex
	c.receiveMutex.Lock()
	defer c.receiveMutex.Unlock()

	// set new timeout
	c.readTimeout = timeout

	// apply new timeout immediately
	_ = c.resetTimeout()
}

// SetMaxWriteDelay will set the maximum amount of time allowed to pass until
// an asynchronous write is flushed.
func (c *BaseConn) SetMaxWriteDelay(delay time.Duration) {
	c.stream.SetMaxWriteDelay(delay)
}

func (c *BaseConn) resetTimeout() error {
	// check timeout
	if c.readTimeout > 0 {
		return c.carrier.SetReadDeadline(time.Now().Add(c.readTimeout))
	}

	return c.carrier.SetReadDeadline(time.Time{})
}
