package main

import (
    "fmt"
)

type N int

func (n N) toString() string {
    return fmt.Sprintf("%#x", n)
}

func (n N) test() {
    fmt.Printf("test.n: %p, %d\n", &n, n)
}

func main() {
    var a N = 25
    fmt.Println(a.toString())
    // receiver类型
    // 无需修改状态的小对象和固定值，用T
    // 引用类型、字符串、函等指针包装对象，用T

    // method expression
    // 会被还原成普通函数样式
    var n N = 1024
    f1 := N.test
    f1(n)

    f2 := (*N).test
    f2(&n)

    // method value
    //  实例或者指针引用的，参数签名不会变，依旧按照正常方式调用
    p := &n
    n++
    f3 := n.test

    n++
    f4 := p.test

    n++
    fmt.Printf("main.n: %p, %v\n", p, n)

    f3()
    f4()
}
