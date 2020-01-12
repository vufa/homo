package transport

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"

	"gopkg.in/tomb.v2"
)

// The WebSocketServer accepts websocket.Conn based connections.
type WebSocketServer struct {
	listener net.Listener
	upgrader *WebSocketUpgrader
	incoming chan *WebSocketConn
	tomb     tomb.Tomb
}

// NewWebSocketServer wraps the provided listener.
func NewWebSocketServer(listener net.Listener, fallback http.Handler) *WebSocketServer {
	// create server
	ws := &WebSocketServer{
		listener: listener,
		upgrader: NewWebSocketUpgrader(fallback),
		incoming: make(chan *WebSocketConn),
	}

	// serve http traffic in background
	ws.tomb.Go(func() error {
		return http.Serve(ws.listener, http.HandlerFunc(ws.handler))
	})

	return ws
}

// CreateWebSocketServer creates a new WS server that listens on the provided address.
func CreateWebSocketServer(address string, fallback http.Handler) (*WebSocketServer, error) {
	// create listener
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	return NewWebSocketServer(listener, fallback), nil
}

// CreateSecureWebSocketServer creates a new WSS server that listens on the
// provided address.
func CreateSecureWebSocketServer(address string, config *tls.Config, fallback http.Handler) (*WebSocketServer, error) {
	// create listener
	listener, err := tls.Listen("tcp", address, config)
	if err != nil {
		return nil, err
	}

	return NewWebSocketServer(listener, fallback), nil
}

func (s *WebSocketServer) handler(w http.ResponseWriter, r *http.Request) {
	// upgrade connection
	conn, _ := s.upgrader.Upgrade(w, r)
	if conn == nil {
		return
	}

	// forward to accept
	select {
	case s.incoming <- conn:
	case <-s.tomb.Dying():
		_ = conn.Close()
	}
}

// Accept will return the next available connection or block until a
// connection becomes available, otherwise returns an error.
func (s *WebSocketServer) Accept() (Conn, error) {
	// await next connection
	select {
	case conn := <-s.incoming:
		return conn, nil
	case <-s.tomb.Dying():
		return nil, s.tomb.Err()
	}
}

// Close will close the underlying listener and cleanup resources. It will
// return an error if the underlying listener didn't close cleanly.
func (s *WebSocketServer) Close() error {
	// kill tomb
	s.tomb.Kill(fmt.Errorf("closed"))

	// close listener
	err := s.listener.Close()
	if err != nil {
		return err
	}

	return nil
}

// Addr returns the server's network address.
func (s *WebSocketServer) Addr() net.Addr {
	return s.listener.Addr()
}

// Upgrader returns the used WebSocketUpgrader.
func (s *WebSocketServer) Upgrader() *WebSocketUpgrader {
	return s.upgrader
}
