package transport

import (
	"crypto/tls"
	"net/http"
	"net/url"
)

// LaunchConfig is used to configure a launcher.
type LaunchConfig struct {
	// The TLS config to be used with secure servers.
	TLSConfig *tls.Config

	// The fallback to be used id a request is not a web socket upgrade.
	WebSocketFallback http.Handler
}

// The Launcher helps with launching a server and accepting connections.
type Launcher struct {
	config LaunchConfig
}

// NewLauncher returns a new Launcher.
func NewLauncher(config LaunchConfig) *Launcher {
	return &Launcher{
		config: config,
	}
}

var sharedLauncher = NewLauncher(LaunchConfig{})

// Launch is a shorthand function.
func Launch(address string) (Server, error) {
	return sharedLauncher.Launch(address)
}

// Launch will launch a server based on information extracted from the address.
func (l *Launcher) Launch(address string) (Server, error) {
	// parse address
	addr, err := url.ParseRequestURI(address)
	if err != nil {
		return nil, err
	}

	// check scheme
	switch addr.Scheme {
	case "tcp", "mqtt":
		return CreateNetServer(addr.Host)
	case "tls", "ssl", "mqtts":
		return CreateSecureNetServer(addr.Host, l.config.TLSConfig)
	case "ws":
		return CreateWebSocketServer(addr.Host, l.config.WebSocketFallback)
	case "wss":
		return CreateSecureWebSocketServer(addr.Host, l.config.TLSConfig, l.config.WebSocketFallback)
	default:
		return nil, ErrUnsupportedProtocol
	}
}
