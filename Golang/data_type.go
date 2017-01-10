package main

import (
    "fmt"
)

func main() {
    // string 支持多种类型编码字符
    // 默认值""
    // 支持!= == < > + +=
    s := "世界\x64\146\u0041"
    // 使用`定义不做转义处理的原始字符串
    ss := `line\r\n line2`
    fmt.Println(s[1], s)
    // byte
    for i := 0; i < len(s); i++ {
        fmt.Printf("%d: [%c]\n", i, s[i])
    }
    // rune
    for i, rr := range s {
        fmt.Printf("%d, [%c]\n", i, rr)
    }
    // 切片
    ss1 := ss[1:4]
    fmt.Println(ss)
    fmt.Println(ss1)
}
