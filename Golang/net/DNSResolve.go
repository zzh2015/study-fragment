package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		fmt.Println("Usage: ", os.Args[0], "hostname")
		os.Exit(1)
	}

	hostname := os.Args[1]

	addr, err := net.ResolveIPAddr("ip", hostname)
	if err != nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}

	fmt.Println("Resolved address is", addr.String())

	addrs, err := net.LookupHost(hostname)
	if err != nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}

	for _, s := range addrs {
		fmt.Println(s)
	}
	os.Exit(0)
}
