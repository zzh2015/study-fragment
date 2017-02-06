package main

import (
    "net"
    "os"
    "fmt"
)

// TCPConn用来作为客户端和服务器之间的通道
// func (c *TCPConn) Write(b []byte) (n int, err os.Error)
// func (c *TCPConn) Read(b []byte) (n int, err os.Error)
// TCP地址信息
/*
    type TCPAddr struct {
        IP IP
        Port int
    }
*/
// func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error) /*net:tcp4,tcp6, tcp*/
// Go通过net包中的DialTCP函数来建立一个TCP连接
// func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Error)
// // laddr为本机地址,raddr为服务地址

// 监听端口,接收客户端请求
// func ListenTCP(net string, laddr *TCPAddr) (l *TCPListener, err os.Error)
// func (l *TCPListener) Accept() (c Conn, err os.Error)

// 设置连接的超时时间
// func DialTimeout(net, addr string, timeout time.Duration) (Conn, error)
// 设置读取/写入一个连接的超时时间
// func (c *TCPConn) SetRaedDeadline(t time.Time) error
// func (c *TCPConn) SetWriteDeadline(t time.Time) error
// func (c *TCPConn) SetKeepAlive(keepalive bool) os.Error
func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s  ip-addr\n", os.Args[0])
        os.Exit(1)
    }

    name := os.Args[1]
    addr := net.ParseIP(name)
    if addr == nil {
        fmt.Println("Invalid address")
    } else {
        fmt.Println("The addr is", addr.String())
    }
    os.Exit(0)
}
