package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(3)

    a, b := make(chan int), make(chan int)

    go func() {
        defer wg.Done()

        for {
            select {
            case x, ok := <-a:
                if !ok {
                    a = nil
                    break
                }
                fmt.Println("value a:", x)
            case x, ok := <-b:
                if !ok {
                    b = nil
                    break
                }
                fmt.Println("value b:", x)
            }
            if a == nil && b == nil {
                return
            }
        }
    }()

    go func() {
        defer wg.Done()
        defer close(a)

        for i := 0; i < 3; i++ {
            a <- i
        }
    }()

    go func() {
        defer wg.Done()
        defer close(b)

        for i := 0; i < 3; i++ {
            b <- i * 10
        }
    }()
    wg.Wait()
}
