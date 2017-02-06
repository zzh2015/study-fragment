package main

import (
    "fmt"
)
// -ldflags "-s" 忽略debug打印信息
// -gcflags "-N -l" 忽略go内部做的一些优化
func main() {
    fmt.Println("use gdb debug do code.")
}
