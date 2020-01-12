// Package transport implements functionality for handling MQTT connections.
package transport

import "errors"

// ErrUnsupportedProtocol is returned if either the launcher or dialer
// couldn't infer the protocol from the URL.
var ErrUnsupportedProtocol = errors.New("unsupported protocol")
