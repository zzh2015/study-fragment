package main

/*import (
    "fmt"
)*/

type tester interface {
    test()
    string() string
}

type data struct {}

func (*data) test() {}
func (data) string() string {return ""}

// 匿名接口类型
type node struct {
    data interface {
        string() string
    }
}

func main() {
    var d data
    // *data实现所有方法集
    var t tester = &d
    t.test()
    t.string()
    // 匿名接口变量
    var tt interface {
        string() string
    } = data{}

    n := node{
        data: tt,
    }

    println(n.data.string())
}
