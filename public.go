// SPDX-License-Identifier: MIT-0
// LICENSE: https://spdx.org/licenses/MIT-0.html

package myip

import (
	"net"
	"context"
)

const (
	// https://unix.stackexchange.com/a/495954/48033

	// NameServer .
	NameServer string = "ns1.google.com:53"
	// SelfAddress .
	SelfAddress string = "o-o.myaddr.l.google.com"
)

// PublicIPs .
func PublicIPs() (ips []net.IP, err error) {
	ip4, err4 := PublicIPv4()
	ip6, err6 := PublicIPv6()
	if err4 != nil {
		err = err4
	} else if err6 != nil {
		err = err6
	}
	return []net.IP{ip4,ip6}, err
}

// PublicIPv4 .
func PublicIPv4() (ip net.IP, err error) {
	return doTXTLookup("udp4")
}

// PublicIPv6 .
func PublicIPv6() (ip net.IP, err error) {
	return doTXTLookup("udp6")
}

func doTXTLookup(transport string) (ip net.IP, err error) {
	// https://unix.stackexchange.com/a/495954/48033
	dns := net.Resolver{
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, transport, NameServer)
		},
	}
	txt, err := dns.LookupTXT(context.Background(), SelfAddress)
	if err != nil {
		return []byte{}, err
	}
	return net.ParseIP(txt[0]), nil
}