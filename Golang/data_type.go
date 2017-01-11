package main

import (
    "fmt"
    "errors"
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
    arr5 := [...]int{10, 3: 100}
    fmt.Println(arr1, arr2, arr3, arr4, arr5)

    // 切片本身只是一个只读对象
    // 可基于数组或者数组指针创建切片,右半开区间
    // len = high - low,可读的写元素数量
    // cap = max - low,数组片段真实长度
    arrs := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    sli:= arrs[2:5]
    for i, value := range sli {
        fmt.Println(i, value)
    }
    //
    sli1 := make([]int, 3, 5)
    sli2 := make([]int, 3)
    sli3 := []int{10, 20, 5: 30}

    fmt.Println(sli1, len(sli1), cap(sli1))
    fmt.Println(sli2, len(sli2), cap(sli2))
    fmt.Println(sli3, len(sli3), cap(sli3))
    // slice achieve stack
    stack := make([]int, 0, 5)
    push := func(x int) error {
        n := len(stack)
        if n == cap(stack) {
            return errors.New("stack is full")
        }

        stack = stack[:n+1]
        stack[n] = x
        return nil
    }

    pop := func() (int, error) {
        n := len(stack)
        if n == 0 {
            return 0, errors.New("stack is empty")
        }

        x := stack[n-1]
        stack = stack[:n-1]

        return x, nil
    }
    // push
    for i := 0; i < 7; i++ {
        fmt.Printf("push %d: %v, %v\n", i, push(i), stack)
    }

    // pop
    for i := 0; i < 7; i++ {
        x, err := pop()
        fmt.Printf("pop: %d, %v, %v\n", x, err, stack)
    }
    // 使用append向slice尾部添加数据

    // 字典
    m := make(map[string]int)
    m["a"] = 1
    m["b"] = 2

    m2 := map[int]struct {
        x int
    } {
        1: {x: 100},
        2: {x: 200},
    }
    fmt.Println(m, m2)
    // ok-idiom判断key是否存在,返回值
    if v, ok := m["b"]; ok {
        fmt.Println(v)
    }

    //// 结构struct
    type node struct {
        _ int
        id int
        next *node
    }

    n1 := node {
        id: 1,
    }

    n2 := node {
        id: 2,
        next: &n1,
    }
    fmt.Println(n1, n2)
    // 空结构作为通道元素类型，用于事件通知
    exit := make(chan struct{})

    go func() {
        fmt.Println("hello, wolrd!")
        exit <- struct{} {}
    } ()

    <-exit
    println(".end")
}
