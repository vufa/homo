package transport

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// The WebSocketUpgrader upgrades HTTP requests to WebSocket connections.
type WebSocketUpgrader struct {
	fallback http.Handler
	upgrader *websocket.Upgrader
	mutex    sync.Mutex
}

// NewWebSocketUpgrader creates a new upgrader with the provided optional
// fallback.
func NewWebSocketUpgrader(fallback http.Handler) *WebSocketUpgrader {
	return &WebSocketUpgrader{
		fallback: fallback,
		upgrader: &websocket.Upgrader{
			HandshakeTimeout:  60 * time.Second,
			ReadBufferSize:    0,
			WriteBufferSize:   0,
			Subprotocols:      []string{"mqtt", "mqttv3.1"},
			Error:             nil,
			CheckOrigin:       func(*http.Request) bool { return true },
			EnableCompression: false,
		},
		mutex: sync.Mutex{},
	}
}

// Upgrade will attempt to upgrade the request and return the connection. If
// the request is not a upgrade it will use the fallback handler if available.
// Encountered errors are already written to the client.
func (u *WebSocketUpgrader) Upgrade(w http.ResponseWriter, r *http.Request) (*WebSocketConn, error) {
	// call fallback if request is not an upgrade
	if r.Header.Get("Upgrade") != "websocket" && u.fallback != nil {
		u.fallback.ServeHTTP(w, r)
		return nil, nil
	}

	// upgrade request
	conn, err := u.upgrader.Upgrade(w, r, nil)
	if err != nil {
		// upgrader already responded to request
		return nil, err
	}

	// create connection
	webSocketConn := NewWebSocketConn(conn)

	return webSocketConn, nil
}

// UnderlyingUpgrader returns the underlying websocket.Upgrader.
func (u *WebSocketUpgrader) UnderlyingUpgrader() *websocket.Upgrader {
	return u.upgrader
}
