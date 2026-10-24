package network

import (
	"fmt"
	"net"
)

func ParseAddress(ipNet *net.IPNet) (Address, error) {
	if ipNet == nil {
		return Address{}, fmt.Errorf("address can`t be nil")
	}

	prefix, bits := ipNet.Mask.Size()
	if bits == 0 {
		return Address{}, fmt.Errorf("incorrect mask network")
	}

	networkIP := ipNet.IP.Mask(ipNet.Mask)

	address := Address{
		IP:      ipNet.IP.String(),
		Mask:    net.IP(ipNet.Mask).String(),
		Prefix:  prefix,
		Network: networkIP.String(),
	}

	return address, nil
}
