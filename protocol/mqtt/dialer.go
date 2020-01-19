package mqtt

import (
	"github.com/256dpi/gomqtt/transport"
	"github.com/countstarlight/homo/utils"
	"net/url"
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

// Dial initiates a connection based in information extracted from an URL.
func (d *Dialer) Dial(urlString string) (transport.Conn, error) {
	uri, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, err
	}
	if uri.Scheme == "ssl" {
		uri.Scheme = "tls"
	}
	return d.Dialer.Dial(uri.String())
}
