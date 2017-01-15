package main

import (
    "time"
)

var c int
func counter() int {
    c++
    return c
}

func main() {
    // go关键字定义并发任务
    go println("p: hello wolrd")
    go func(s string) {
        println("f:", s)
    }("hello world")
    println("main:hello")

    a := 100
    go func(x, y int) {
        time.Sleep(time.Second)
        println("go: ", x, y)
    } (a, counter())

    a += 100
    println("main: ", a, counter())
    time.Sleep(time.Second*3)
}
