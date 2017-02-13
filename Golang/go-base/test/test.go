package main

import (
    "./lib"
    "fmt"
    "unsafe"
)
// 绕过访问权限
// go run test.go
func main() {
    d := lib.NewData()
    d.Y = 200

    p := (*struct{x int})(unsafe.Pointer(d))
    p.x = 100

    fmt.Printf("%+v\n", *d)
}
