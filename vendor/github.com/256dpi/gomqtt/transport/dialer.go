package transport

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

// DialConfig is used to configure a dialer.
type DialConfig struct {
	// The TLS config to be used with secure connections.
	TLSConfig *tls.Config

	// The additional request headers for web socket connections.
	RequestHeader http.Header

	// The time after which a dial attempt is cancelled.
	//
	// Default: No timeout.
	Timeout time.Duration

	// The default ports to be uses if no port has been specified.
	//
	// Defaults: 1883, 8883, 80, 443.
	DefaultTCPPort string
	DefaultTLSPort string
	DefaultWSPort  string
	DefaultWSSPort string
}

func (c *DialConfig) ensureDefaults() {
	// set default ports
	if c.DefaultTCPPort == "" {
		c.DefaultTCPPort = "1883"
	}
	if c.DefaultTLSPort == "" {
		c.DefaultTLSPort = "8883"
	}
	if c.DefaultWSPort == "" {
		c.DefaultWSPort = "80"
	}
	if c.DefaultWSPort == "" {
		c.DefaultWSPort = "443"
	}
}

// The Dialer handles connecting to a server and creating a connection.
type Dialer struct {
	config    DialConfig
	netDialer net.Dialer
	wsDialer  websocket.Dialer
}

// NewDialer returns a new Dialer.
func NewDialer(config DialConfig) *Dialer {
	// ensure defaults
	config.ensureDefaults()

	return &Dialer{
		config: config,
		netDialer: net.Dialer{
			Timeout: config.Timeout,
		},
		wsDialer: websocket.Dialer{
			Proxy:            http.ProxyFromEnvironment,
			TLSClientConfig:  config.TLSConfig,
			HandshakeTimeout: config.Timeout,
			Subprotocols:     []string{"mqtt"},
		},
	}
}

var sharedDialer = NewDialer(DialConfig{})

// Dial is a shorthand function.
func Dial(address string) (Conn, error) {
	return sharedDialer.Dial(address)
}

// Dial initiates a connection based in information extracted from an URL.
func (d *Dialer) Dial(address string) (Conn, error) {
	// parse address
	addr, err := url.ParseRequestURI(address)
	if err != nil {
		return nil, err
	}

	// get host and port
	host, port, err := net.SplitHostPort(addr.Host)
	if err != nil {
		host = addr.Host
		port = ""
	}

	// check scheme
	switch addr.Scheme {
	case "tcp", "mqtt":
		// set default port
		if port == "" {
			port = d.config.DefaultTCPPort
		}

		// make connection
		conn, err := d.netDialer.Dial("tcp", net.JoinHostPort(host, port))
		if err != nil {
			return nil, err
		}

		return NewNetConn(conn), nil
	case "tls", "ssl", "mqtts":
		// set default port
		if port == "" {
			port = d.config.DefaultTLSPort
		}

		// make connection
		conn, err := tls.DialWithDialer(&d.netDialer, "tcp", net.JoinHostPort(host, port), d.config.TLSConfig)
		if err != nil {
			return nil, err
		}

		return NewNetConn(conn), nil
	case "ws":
		// set default port
		if port == "" {
			port = d.config.DefaultWSPort
		}

		// format url
		wsURL := fmt.Sprintf("ws://%s:%s%s", host, port, addr.Path)

		// make connection
		conn, _, err := d.wsDialer.Dial(wsURL, d.config.RequestHeader)
		if err != nil {
			return nil, err
		}

		return NewWebSocketConn(conn), nil
	case "wss":
		// set default port
		if port == "" {
			port = d.config.DefaultWSSPort
		}

		// format url
		wsURL := fmt.Sprintf("wss://%s:%s%s", host, port, addr.Path)

		// make connection
		conn, _, err := d.wsDialer.Dial(wsURL, d.config.RequestHeader)
		if err != nil {
			return nil, err
		}

		return NewWebSocketConn(conn), nil
	default:
		return nil, ErrUnsupportedProtocol
	}
}
