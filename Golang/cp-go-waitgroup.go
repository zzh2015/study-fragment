package main

import (
    "fmt"
    "time"
    // "runtime"
    "sync"
)

func main() {
    var wg sync.WaitGroup

    for i:= 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()

            time.Sleep(time.Second)
            fmt.Println("goroutine", id, "done")
        }(i)
    }

    // runtime.GOMAXPROCS(3) 设置并发线程数
    fmt.Println("main...")
    wg.Wait()
    fmt.Println("main exit.")
}
