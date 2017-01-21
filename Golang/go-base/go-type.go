package main

import (
    "fmt"
    "strconv"
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

    //
    gb, goo, gh := 100, 0144, 0x64
    fmt.Printf("0b%b, %#o, %#x\n", gb, goo, gh)

    gx, _ := strconv.ParseInt("1100100", 2, 32)
    gy, _ := strconv.ParseInt("0144", 8, 32)
    gz, _ := strconv.ParseInt("64", 16, 32)
    fmt.Println(gx, gy, gz)

    println("0b" + strconv.FormatInt(gx, 2))
    println("0" + strconv.FormatInt(gx, 8))
    println("0x" + strconv.FormatInt(gx, 16))
}
