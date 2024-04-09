package main

import(
	"fmt"
	"strings"
	"strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	toStr := make([]string, len(ip))
	for i, val := range ip {
		toStr[i] = strconv.Itoa(int(val))
	}	
	return strings.Join(toStr, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
