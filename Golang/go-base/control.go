package main

import (
    // . "fmt" 省略前缀包名
    // f "fmt" 别名
    // _ "github.com/zzh/mysql/godrv" 引入该包，使用包里的init函数
    "fmt"
)

func main() {
    // if xinit();a < b {
    // }
    // a, b作用域为条件代码块内
    if a, b := 1024, 2045; a < b {
        fmt.Println("control flow--if")
    }

    // switch
    u, v, w, z := 2, 3, 5, 3
    switch z {
    case u, v:
        fmt.Println("match u or v")
        // falltrought
    case w:
        fmt.Println("match w")
    default:
        fmt.Println("default")
    }

    switch jf := false; {
    case jf:
        fmt.Println("use switch replace if")
    case !jf:
        fmt.Println("true")
    }

    for index := 0; index < 3; index++ {
        fmt.Println("index=", index)
    }

    count := 0
    for count < 3 {
        fmt.Println("count=", count)
        count++
    }

    user := [3]string{"bob", "tom", "amy"}
    for index, data := range user {
        fmt.Println("index=", index, "user=", data)
    }
}
