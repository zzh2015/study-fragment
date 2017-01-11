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

    // 数组
    // 定义多维数组，只有第一维允许使用[...]
    // len,cap返回第一维度长度
    // 复制，会复制整个数组
    var arr1 [5]int // zero
    arr2 := [4]int{2, 5}
    arr3 := [4]int{5, 3: 10}
    arr4 := [...]int{1, 2, 3}
    arr5 := [...]int{10; 3: 100}
    fmt.Println(arir1, arr2, arr3, arr4, arr5)

    // 切片本身只是一个只读对象
}
