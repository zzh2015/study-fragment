package main

import (
    "fmt"
)

func main() {
    var x int   // auto init zero value x = 0
    var y = false //auto type
    var a, s = 100, "hello"

    if y {
        fmt.Println("Hello world!")
    } else {
        fmt.Println("x=", x, " a=", a, " s=", s)
    }

    // short simple define
    // only in func internal
    xx, yy := 1024, "bye"
    fmt.Println("xx=", xx, " yy=", yy)

    // x =, ss define
    x, ss := 2048, "good night"
    fmt.Println("x=", x, "ss=", ss)
}
