package discover

// TODO: decide out how, when, and if to return the Floating IP (URLs noted below)

import (
	"net"
)

const (
	privatePublicIPv4URL = "http://192.168.0.210?format=json"
	privatePublicIPv6URL = "http://192.168.0.210?format=json"
)

// NewDigitalOceanDiscoverer returns a new Digital Ocean network discoverer
func NewPrivateDiscoverer() Discoverer {
	return NewDiscoverer(
		PublicIPv4DiscovererOption(privatePublicIPv4),
		PublicIPv6DiscovererOption(privatePublicIPv6),
	)
}

func privatePublicIPv4() (net.IP, error) {
	return JsonIPFromHTTP(privatePublicIPv4URL, nil)
}

func privatePublicIPv6() (net.IP, error) {
	return JsonIPFromHTTP(privatePublicIPv6URL, nil)
}
