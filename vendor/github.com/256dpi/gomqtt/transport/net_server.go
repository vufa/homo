package transport

import (
	"crypto/tls"
	"net"
)

// A NetServer accepts net.Conn based connections.
type NetServer struct {
	listener net.Listener
}

// NewNetServer wraps the provided listener.
func NewNetServer(listener net.Listener) *NetServer {
	return &NetServer{
		listener: listener,
	}
}

// CreateNetServer creates a new TCP server that listens on the provided address.
func CreateNetServer(address string) (*NetServer, error) {
	// create listener
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	return NewNetServer(listener), nil
}

// CreateSecureNetServer creates a new TLS server that listens on the provided address.
func CreateSecureNetServer(address string, config *tls.Config) (*NetServer, error) {
	// create listener
	listener, err := tls.Listen("tcp", address, config)
	if err != nil {
		return nil, err
	}

	return NewNetServer(listener), nil
}

// Accept will return the next available connection or block until a
// connection becomes available, otherwise returns an error.
func (s *NetServer) Accept() (Conn, error) {
	// wait next connection
	conn, err := s.listener.Accept()
	if err != nil {
		return nil, err
	}

	return NewNetConn(conn), nil
}

// Close will close the underlying listener and cleanup resources. It will
// return an error if the underlying listener didn't close cleanly.
func (s *NetServer) Close() error {
	return s.listener.Close()
}

// Addr returns the server's network address.
func (s *NetServer) Addr() net.Addr {
	return s.listener.Addr()
}
