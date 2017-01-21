package main

import (
    "fmt"
    "sync"
)
// factory mode
type receiver struct {
    sync.WaitGroup
    data chan int
}

func newreceiver() *receiver {
    r := &receiver {
        data: make(chan int),
    }

    r.Add(1)
    go func() {
        defer r.Done()
        for x := range r.data {
            fmt.Println("r value:", x)
        }
    }()

    return r
}

func main() {
    r := newreceiver()

    r.data <-1
    r.data <-2

    close(r.data)
    r.Wait()
}
