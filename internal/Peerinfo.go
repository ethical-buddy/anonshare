



package internal

import (
	
	"net"
	"log"
)

// GetLocalIP returns the first non-loopback IPv4 address as a string
func GetLocalIP() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal("Error getting network interfaces:", err)	
		
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			log.Printf("Error getting addresses for interface %s: %v", iface.Name, err)
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue
			}

			return ip.String() 
		}
	}
	log.Fatal("Could not detect a non-loopback IPv4 address")
	return "" 
}




