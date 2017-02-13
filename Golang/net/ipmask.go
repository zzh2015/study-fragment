package main

import (
    "fmt"
    "os"
    "net"
)

// type IPMask[] byte
func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr\n", os.Args[0])
        os.Exit(1)
    }
    dotAddr := os.Args[1]

    // get ip
    addr := net.ParseIP(dotAddr)
    if addr == nil {
        fmt.Println("Invalid address")
        os.Exit(1)
    }

    // get mask
    mask := addr.DefaultMask()
    network := addr.Mask(mask)
    ones, bits := mask.Size()
    fmt.Println("Address is ", addr.String(),
        "Default mask len is ", bits,
        "Leading ones count is ", ones,
        "Mask is ", mask.String(),
        "NetWork is ", network.String())
    os.Exit(0)
}
