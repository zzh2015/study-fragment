package main

import (
    "fmt"
)

func main() {
    // 使用ok-idom或range模式处理数据
    done := make(chan struct{})
    ch := make(chan int)

    go func() {
        defer close(done)

        /*for {
            x, ok := <-ch
            if !ok {
                return
            }

            fmt.Println("value:", x)
        }*/
        for x := range ch {
            fmt.Println("value:", x)
        }
    }()

    ch <- 1024
    ch <- 2048
    ch <- 4096

    close(ch)

    <-done
}
