package main

import "fmt"

type IPAddr [4]byte

func (ipaddr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3])
}

func (ipaddr IPAddr) String1() string {
	s := ""
	for i := 0; i < 4; i++ {
		s += string(ipaddr[i]) + "."
	}
	return s
}

// TODO: Add a "String() string" method to IPAddr.
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
