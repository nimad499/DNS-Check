package main

import (
	"testing"
)

func TestValidateDNS(t *testing.T) {
	tests := []struct {
		dnsServer string
		expected  bool
	}{
		{"192.168.1.1", true},
		{"255.255.255.255", true},
		{"0.0.0.0", true},
		{"256.256.256.256", false},
		{"192.168.1", false},
		{"192.168.1.300", false},
		{"abc.def.ghi.jkl", false},
		{"123.456.78.90", false},
		{"192.168.1.01", true},
		{"192.168.01.1", true},
	}

	for _, tc := range tests {
		t.Run(tc.dnsServer, func(t *testing.T) {
			result := ValidateDNS(tc.dnsServer)
			if result != tc.expected {
				t.Errorf("ValidateDNS(%q) = %v; expected %v", tc.dnsServer, result, tc.expected)
			}
		})
	}
}
