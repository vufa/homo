package mqtt

import (
	"github.com/256dpi/gomqtt/transport"
	"github.com/countstarlight/homo/utils"
)

// The Dialer handles connecting to a server and creating a connection.
type Dialer struct {
	*transport.Dialer
}

// NewDialer returns a new Dialer.
func NewDialer(c utils.Certificate) (*Dialer, error) {
	tls, err := utils.NewTLSClientConfig(c)
	if err != nil {
		return nil, err
	}
	d := &Dialer{Dialer: transport.NewDialer(transport.DialConfig{TLSConfig: tls})}
	return d, nil
}
