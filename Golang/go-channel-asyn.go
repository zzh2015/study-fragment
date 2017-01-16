package main

import (
    "fmt"
    "unsafe"
)

func main() {
    // 同步模式必须有配对的goroutine出现,否则会一直阻塞
    // 异步模式在缓冲区未满或者数据未读完前,不会阻塞
    data := make(chan int, 3) //创建3个带缓冲槽的异步通道

    data <- 1024
    data <- 2048

    fmt.Println(<-data)
    fmt.Println(<-data)
    // 通常异步通道，异步通道有助于提升性能，减少排队阻塞

    // 通道变量本身是指针
    // 可用操作符判断是否为同一对象或nil
    var a, b chan int = make(chan int, 3), make(chan int)
    var c chan bool

    fmt.Println(a == b)
    fmt.Println(c == nil)

    fmt.Printf("%p, %d\n", a, unsafe.Sizeof(a))

    // cap和len返回缓冲区的大小和当前已缓冲数量,对于同步通道则都返回0
    a <-1
    a <-2

    fmt.Println("a:", cap(a), len(a))
    fmt.Println("b:", cap(b), len(b))
}
