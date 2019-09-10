package discover

// TODO: decide out how, when, and if to return the Floating IP (URLs noted below)

import (
	"net"
)

const (
	doPublicIPv4URL = "http://192.168.0.210?format=json"
	doPublicIPv6URL = "http://192.168.0.210?format=json"
)

// NewDigitalOceanDiscoverer returns a new Digital Ocean network discoverer
func NewPrivateDiscoverer() Discoverer {
	return NewDiscoverer(
		PublicIPv4DiscovererOption(doPublicIPv4),
		PublicIPv6DiscovererOption(doPublicIPv6),
	)
}

func doPublicIPv4() (net.IP, error) {
	return StandardIPFromHTTP(doPublicIPv4URL, nil)
}

func doPublicIPv6() (net.IP, error) {
	return StandardIPFromHTTP(doPublicIPv6URL, nil)
}
