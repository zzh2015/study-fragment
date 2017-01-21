package main

import (
    "fmt"
    "time"
)

func main() {
    exit := make(chan struct{})

    go func() {
        time.Sleep(time.Second)
        fmt.Println("goroutine done.")

        // 写入数据也可解除阻塞
        close(exit)
    }()

    println("main...")
    <-exit
    println("main exit.")
}
