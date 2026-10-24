package network

import (
	"net"
	"testing"
)

func TestParseAddress(t *testing.T) {
	tests := []struct {
		name        string
		ipNet       *net.IPNet
		wantIP      string
		wantMask    string
		wantPrefix  int
		wantNetwork string
	}{
		{
			name: "IPv4 /24",
			ipNet: &net.IPNet{
				IP:   net.ParseIP("192.168.10.11"),
				Mask: net.CIDRMask(24, 32),
			},
			wantIP:      "192.168.10.11",
			wantMask:    "255.255.255.0",
			wantPrefix:  24,
			wantNetwork: "192.168.10.0",
		},
		{
			name: "IPv4 /16",
			ipNet: &net.IPNet{
				IP:   net.ParseIP("169.254.88.59"),
				Mask: net.CIDRMask(16, 32),
			},
			wantIP:      "169.254.88.59",
			wantMask:    "255.255.0.0",
			wantPrefix:  16,
			wantNetwork: "169.254.0.0",
		},
		{
			name: "IPv6 /64",
			ipNet: &net.IPNet{
				IP:   net.ParseIP("2001:db8:abcd:12::1234"),
				Mask: net.CIDRMask(64, 128),
			},
			wantIP:      "2001:db8:abcd:12::1234",
			wantMask:    "ffff:ffff:ffff:ffff::",
			wantPrefix:  64,
			wantNetwork: "2001:db8:abcd:12::",
		},
		{
			name: "IPv4 /32",
			ipNet: &net.IPNet{
				IP:   net.ParseIP("192.168.1.10"),
				Mask: net.CIDRMask(32, 32),
			},
			wantIP:      "192.168.1.10",
			wantMask:    "255.255.255.255",
			wantPrefix:  32,
			wantNetwork: "192.168.1.10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAddress(tt.ipNet)
			if err != nil {
				t.Fatalf("ParseAddress() вернула неожиданную ошибку: %v", err)
			}

			if got.IP != tt.wantIP {
				t.Errorf("IP = %q, ожидалось %q", got.IP, tt.wantIP)
			}

			if got.Mask != tt.wantMask {
				t.Errorf("Mask = %q, ожидалось %q", got.Mask, tt.wantMask)
			}

			if got.Prefix != tt.wantPrefix {
				t.Errorf("Prefix = %d, ожидалось %d", got.Prefix, tt.wantPrefix)
			}

			if got.Network != tt.wantNetwork {
				t.Errorf("Network = %q, ожидалось %q", got.Network, tt.wantNetwork)
			}
		})
	}
}

func TestParseAddressNil(t *testing.T) {
	_, err := ParseAddress(nil)

	if err == nil {
		t.Fatal("ожидалась ошибка для nil-адреса")
	}
}

func TestParseAddressInvalidMask(t *testing.T) {
	ipNet := &net.IPNet{
		IP:   net.ParseIP("192.168.1.10"),
		Mask: net.IPMask{255, 0, 255, 0},
	}

	_, err := ParseAddress(ipNet)

	if err == nil {
		t.Fatal("ожидалась ошибка для некорректной маски")
	}
}
