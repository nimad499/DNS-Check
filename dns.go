package main

import (
	"context"
	"net"
	"net/http"
	"regexp"
	"time"
)

type DNS struct {
	name      string
	dnsServer string
}

func customResolver(dnsServer string) *net.Resolver {
	return &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return net.DialTimeout(network, dnsServer+":53", 5*time.Second)
		},
	}
}

func customDialContext(dnsServer string) func(ctx context.Context, network, addr string) (net.Conn, error) {
	resolver := customResolver(dnsServer)
	dialer := net.Dialer{
		Resolver: resolver,
	}

	return dialer.DialContext
}

func sendRequestWithCustomDNS(host string, dnsServer string) (int, error) {
	transport := &http.Transport{
		DialContext: customDialContext(dnsServer),
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   5 * time.Second,
	}

	resp, err := client.Head(host)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

func resolveDNS(host string, dnsServer string) ([]string, error) {
	resolver := customResolver(dnsServer)

	return resolver.LookupHost(context.Background(), host)
}

func (dns *DNS) Check(host string) (int, error) {
	if addrs, err := resolveDNS(host, dns.dnsServer); err == nil && len(addrs) != 0 {
		status, err := sendRequestWithCustomDNS("https://"+host, dns.dnsServer)

		return status, err
	} else {
		return 0, err
	}
}

func ValidateDNS(dnsServer string) bool {
	pattern := `^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`
	re := regexp.MustCompile(pattern)

	if re.MatchString(dnsServer) {
		return true
	}
	return false
}
