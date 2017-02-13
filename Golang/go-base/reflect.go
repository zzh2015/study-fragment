package main

import (
    "fmt"
    "reflect"
)

// reflect
// struct cast to pointer,then iterate it

type user struct {
    name string
    age int
}

type manager struct {
    user
    title string
}

func main() {
    var m manager
    t := reflect.TypeOf(&m) // to pointer

    // Kind()底层类型
    if t.Kind() == reflect.Ptr {
        t = t.Elem()
    }

    for i := 0; i < t.NumField(); i++ {
        f := t.Field(i)
        fmt.Println(f.Name, f.Type, f.Offset)
        if f.Anonymous { // 匿名字段
            for x := 0; x < f.Type.NumField(); x++ {
                af := f.Type.Field(x)
                fmt.Println(" ", af.Name, af.Type)
            }
        }
    }
}
