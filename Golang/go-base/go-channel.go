package main

import (
    "fmt"
)

func main() {
    // go使用CSP通道来替代内存共享，实现并发安全
    // 如果消息未能及时处理会阻塞当前端
    // 通道只是一个队列
    // 同步状态下，发送方和接受方配对，然后直接复制数据给对方
    // 异步状态下,抢占数据缓冲槽.发送方有空槽可供写入，接收方有缓冲数据可读.
    fmt.Println("go channel")
    // 除传递消息外,通道还长用作时间通知
    notify := make(chan struct{})
    data := make(chan string)

    go func() {
        s := <-data
        fmt.Println(s)
        close(notify)
    }()

    data <- "send data"
    <-notify
}
