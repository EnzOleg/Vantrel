package main

import (
	"fmt"
	"log"

	"github.com/EnzOleg/vantrel/internal/network"
)

func main() {
	interfaces, err := network.Interfaces()
	if err != nil {
		log.Fatalf("Не удалось получить сетевые интерфейсы: %v", err)
	}
	for _, iface := range interfaces {
		fmt.Printf("\nИнтерфейс: %s\n", iface.Name)
		fmt.Printf("	MTU: %d\n", iface.MTU)
		fmt.Printf("	MAC: %s\n", iface.MAC)
		fmt.Printf("	Флаги: %s\n", iface.Flags)

		for _, addr := range iface.Addresses {
			fmt.Printf("	IP: %s | Маска: %s\n", addr.IP, addr.Mask)
			fmt.Printf("	Префикс: %d\n", addr.Prefix)
			fmt.Printf("	Подсеть: %s\n", addr.Network)
		}
	}
	fmt.Println("Vantrel v0.2.0")
}
