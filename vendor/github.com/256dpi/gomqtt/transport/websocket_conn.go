package transport

import (
	"errors"
	"io"
	"net"
	"time"

	"github.com/gorilla/websocket"
)

// ErrNotBinary may be returned by WebSocket connection when a message is
// received that is not binary.
var ErrNotBinary = errors.New("received web socket message is not binary")

type wsStream struct {
	conn   *websocket.Conn
	reader io.Reader
}

func (s *wsStream) Read(buf []byte) (int, error) {
	// prepare total
	total := 0

	for {
		// get next reader
		if s.reader == nil {
			messageType, reader, err := s.conn.NextReader()
			if _, ok := err.(*websocket.CloseError); ok {
				return 0, io.EOF
			} else if err != nil {
				return 0, err
			} else if messageType != websocket.BinaryMessage {
				return 0, ErrNotBinary
			}

			// set current reader
			s.reader = reader
		}

		// read data
		n, err := s.reader.Read(buf)

		// increment counter
		total += n
		buf = buf[n:]

		// handle EOF
		if err == io.EOF {
			// clear reader
			s.reader = nil

			continue
		}

		return total, err
	}
}

func (s *wsStream) Write(buf []byte) (n int, err error) {
	// create writer
	writer, err := s.conn.NextWriter(websocket.BinaryMessage)
	if err != nil {
		return 0, err
	}

	// write buffer
	n, err = writer.Write(buf)
	if err != nil {
		return n, err
	}

	// close writer
	err = writer.Close()
	if err != nil {
		return n, err
	}

	return n, nil
}

func (s *wsStream) Close() error {
	// Close can be called during read and write, therefore we cannot write a
	// close message to the client without risking a concurrent write on the
	// websocket conn. The MQTT spec anyway requires clients to terminate the
	// connection, therefore we don't have to really care about announcing a
	// server-side connection close.

	return s.conn.Close()
}

func (s *wsStream) SetReadDeadline(t time.Time) error {
	return s.conn.SetReadDeadline(t)
}

// The WebSocketConn wraps a websocket.Conn. The implementation supports packets
// that are chunked over several WebSocket messages and packets that are coalesced
// to one WebSocket message.
type WebSocketConn struct {
	*BaseConn

	conn *websocket.Conn
}

// NewWebSocketConn returns a new WebSocketConn.
func NewWebSocketConn(conn *websocket.Conn) *WebSocketConn {
	return &WebSocketConn{
		BaseConn: NewBaseConn(&wsStream{conn: conn}),
		conn:     conn,
	}
}

// LocalAddr returns the local network address.
func (c *WebSocketConn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

// RemoteAddr returns the remote network address.
func (c *WebSocketConn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

// UnderlyingConn returns the underlying websocket.Conn.
func (c *WebSocketConn) UnderlyingConn() *websocket.Conn {
	return c.conn
}
