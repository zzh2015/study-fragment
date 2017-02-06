package main

import (
    "fmt"
    "net"
    "os"
    "time"
    "strings"
    "strconv"
)

func main() {
    service := ":7777"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)
    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        // daytime := time.Now().String()
        // conn.Write([]byte(daytime))
        // // 关闭连接
        // conn.Close()

        go handleClient(conn)
    }
}

func handleClient(conn net.Conn) {
    // defer conn.Close()
    // daytime := time.Now().String()
    // conn.Write([]byte(daytime))

    conn.SetReadDeadline(time.Now().Add(time.Minute))
    request := make([]byte, 128)
    defer conn.Close()
    for {
        read_len ,err := conn.Read(request)

        if err != nil {
            fmt.Println(err)
            break
        }

        if read_len == 0 {
            break
        } else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
            daytime := strconv.FormatInt(time.Now().Unix(), 10)
            conn.Write([]byte(daytime))
        } else {
            daytime := time.Now().String()
            conn.Write([]byte(daytime))
        }
        request = make([]byte, 128)
    }
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
