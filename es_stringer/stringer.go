package main

import (
	"fmt"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
// This method, implementing the String method could be
// declaer toward the Stringer interface.
//type Stringer interface {
//    String() string
// The fmt package (and many others) look for this interface to print values.
// Is used to print values.
func (ip IPAddr) String() string {
	// Use String() method of interface Stringer implemented by IPAddr to print.
	// Like an overload of ToString() method.
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
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
