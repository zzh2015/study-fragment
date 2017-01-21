package main

import (
    "fmt"
    "errors"
)

// function feature
// 无需前置声明
// 不支持嵌套命名和函数重载
// 不支持默认参数
// 支持不定长变参
// 支持多返回值
// 支持命名返回值
// 支持匿名函数和闭包
func hello() {
    fmt.Println("hello,世界!")
}

func exec(f func()) {
    f()
}

// 函数返回局部变量是安全的
func get_value() * int {
    x := 1024
    return &x
}
// 变参本质上是一个切片,只能接收一到多个同类型参数
func varArgs(s string, a ...int) {
    fmt.Printf("%T, %v\n", a, a)
}

// 多返回值
func multiRetValue(x, y int) (int, error) {
    if y == 0 {
        return 0, errors.New("division by zero")
    }

    return x/y, nil
}

// 命名返回值
func namingRetValue(sql string, index int) (count int, pages int, err error) {
    if index == 0 {
        count = 0
        pages = 0
        err = errors.New("invalid sql")
        // 隐式返回
        return
    }

    fmt.Println(sql)
    count = 1
    pages = 1
    err = nil
    return

}

// 闭包
func test() []func() {
    var s []func()

    for i:=0; i<2; i++ {
        x := i
        s = append(s, func() {
            fmt.Println(&x, x)
        })
    }

    return s
}

func main() {
    f := hello
    exec(f)
    func_value := get_value() // var func_value *int
    fmt.Println(func_value, *func_value)

    varArgs("abc", 2, 3, 5, 8)
    // 忽略部分返回值
    div, _ := multiRetValue(1024, 2)
    fmt.Println(div)
    //
    count, pages, _ := namingRetValue("select * from vps_info;", 1)
    fmt.Println(count, pages)

    //
    defer fmt.Println("leave main last")
    // 可以用作参数或者返回值
    func(s string) {
        fmt.Println(s)
    } ("匿名函数")

    // 遍历
    for _, tf := range test() {
        tf()
    }

    // defer延迟调用
    // 直到当前函数执行结束前才被执行
    // 常用于资源释放、解除锁定、以及错误处理等作用
    defer fmt.Println("leave main first")
}
