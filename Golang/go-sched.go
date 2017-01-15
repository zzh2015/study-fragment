package main

import (
    "fmt"
    "runtime"
)

func main() {
    runtime.GOMAXPROCS(1)

    exit := make(chan struct{})

    go func() {
        defer close(exit)

        go func() {
            fmt.Printf("b\n")
        }()

        for i := 0; i < 4; i++ {
            fmt.Println("a:", i)
            if i == 1 {
                runtime.Gosched()
            }
        }
    }()
    // Goexit立即终止当前任务 运行时确保所有已注册的延迟调用被执行

    <-exit
}
