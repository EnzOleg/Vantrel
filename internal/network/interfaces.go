package network

import (
	"log"
	"net"
)

type Address struct {
	IP      string
	Mask    string
	Prefix  int
	Network string
}

type NetworkInterface struct {
	Name      string
	MTU       int
	MAC       string
	Flags     string
	Addresses []Address
}

func Interfaces() ([]NetworkInterface, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var result []NetworkInterface

	for _, ifi := range interfaces {
		addrs, err := ifi.Addrs()
		if err != nil {
			log.Printf("Не удалось получить адреса для %s: %v", ifi.Name, err)
			continue
		}

		item := NetworkInterface{
			Name:  ifi.Name,
			MTU:   ifi.MTU,
			MAC:   ifi.HardwareAddr.String(),
			Flags: ifi.Flags.String(),
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			address, err := ParseAddress(ipNet)
			if err != nil {
				log.Printf("Не удалось обработать адрес интерфейса %s: %v", ifi.Name, err)
				continue
			}

			item.Addresses = append(item.Addresses, address)
		}

		result = append(result, item)
	}
	return result, nil
}
