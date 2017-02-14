package main

import (
	"fmt"
	"net"
	"os"
)

// 查询服务对应的端口
// ./lookupport tcp telnet
func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr,
			"Usage: %s network-type service\n",
			os.Args[0])
	}

	networkType := os.Args[1]
	service := os.Args[2]

	port, err := net.LookupPort(networkType, service)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}

	fmt.Println("The Service port is", port)
	os.Exit(0)
}
