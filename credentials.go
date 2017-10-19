package gotelemetry

import (
	"net/url"
)

// Credentials incorporates the information required to call the Telemetry
// service. Normally, you will only need to provide an API token, but you can
// also provide a custom server URL if so required
type Credentials struct {
	APIKey       string     // The API Key
	ServerURL    *url.URL   // The URL should be in the format "http(s)://host/"
	DebugChannel chan error // An optional channel that receives debug messages
}

// NewCredentials function
func NewCredentials(apiKey string, serverURL ...string) (Credentials, error) {
	server := "https://api.telemetrytv.com"

	if len(serverURL) > 0 {
		server = serverURL[0]
	}

	url, err := url.Parse(server)

	return Credentials{apiKey, url, nil}, err
}

// SetDebugChannel function
func (c *Credentials) SetDebugChannel(debugChannel chan error) {
	c.DebugChannel = debugChannel
}
