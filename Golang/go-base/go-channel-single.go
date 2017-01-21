package main

import (
    "fmt"
    "sync"
)

func main() {
    // 使用类型转换来获取单向通道
    var wg sync.WaitGroup
    wg.Add(2)

    c := make(chan int)
    var send chan<- int = c
    var recv <-chan int = c

    go func() {
        defer wg.Done()

        for x := range recv {
            fmt.Println("value:", x)
        }
    }()

    go func() {
        defer wg.Done()
        defer close(c) // close(recv)无效

        for i:= 0; i < 3; i++ {
            send <- i
        }
    }()

    wg.Wait()
}
