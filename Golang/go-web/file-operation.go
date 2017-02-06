package main

import (
    "fmt"
    "os"
)

func main() {
    // create dir
    os.Mkdir("aaa", 0777)
    os.MkdirAll("aaa/bbb", 0777)
    err := os.Remove("aaa")
    if err != nil {
        fmt.Println(err)
    }
    os.RemoveAll("aaa")

    // write
    fileName := "fp.txt"
    fout, err := os.Create(fileName)
    if err != nil {
        fmt.Println(fileName, err)
        return
    }
    defer fout.Close()
    for i := 0; i < 10; i++ {
        fout.WriteString("Hello World\r\n")
        fout.Write([]byte("Hello World\r\n"))
    }

    // read
    fin, err := os.Open(fileName)
    if err != nil {
        fmt.Println(fileName, err)
        return
    }
    defer fin.Close()
    buf := make([]byte, 1024)
    for {
        n, _ := fin.Read(buf)
        if 0 == n {
            break
        }
        os.Stdout.Write(buf[:n])
    }
}
