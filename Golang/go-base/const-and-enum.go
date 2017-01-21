package main

import (
    "fmt"
)

// define const type
const x, y int = 1123, 5813
const s = "const type"
const (
    z int = 10
    ch byte = byte(z) // cast
)

// golang not build inter enum
type follow uint32
const (
    _ follow = iota // iota means auto +1
    KB = 1 << (10 * iota)
    MB = 1 << (10 * iota)
    GB = 1 << (10 * iota)
)

func test_enum(f follow) {
    println(f)
}

func main() {
    fmt.Println("x=", x, "y=", y)
    fmt.Println("s=", s)
    fmt.Println("z=", z, "ch=", ch)
    {
        const str = "happy chris"
        fmt.Println("str=", str)
    }
    test_enum(MB)
}
