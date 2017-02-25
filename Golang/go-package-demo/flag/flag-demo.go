package main

import (
    "flag"
    "fmt"
)

// -flag 只支持bool
// -flag=x
// -flag x 支持非bool
func main() {
    var hostIP, port string
    var level int64
    var enable bool

    flag.BoolVar(&enable, "e", false, "is enable?")
    flag.Int64Var(&level, "l", 0, "desc level")
    flag.StringVar(&hostIP, "s", "", "host ip")
    flag.StringVar(&port, "p", "", "server port")

    flag.Parse()

    if enable {
        fmt.Println("level: ", level)
    }

    fmt.Println("host ip: ", hostIP)
    fmt.Println("server port: ", port)
}
