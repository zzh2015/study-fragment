package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

// a client
// tcp-client x.x.x.x:80
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr,
			"Usage: %s network-type service\n",
			os.Args[0])
	}
	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	// send data
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	// read data
	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
