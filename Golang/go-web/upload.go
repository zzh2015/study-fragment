package main

import (
    "fmt"
    "html/template"
    "net/http"
    "log"
    "time"
    "os"
    "crypto/md5"
    "io"
    "strconv"
)

func upload(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method: ", r.Method)
    if r.Method == "GET" {
        crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
        token := fmt.Sprintf("%x", h.Sum(nil))
        t, _ := template.ParseFiles("upload.gtpl")
        t.Execute(w, token)
    } else {
        r.ParseMultipartForm(32 << 20)
        file, handler, err := r.FormFile("uploadfile")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()
        fmt.Fprintf(w, "%v", handler.Header)
        f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY | os.O_CREATE, 0666)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer f.Close()
        io.Copy(f, file)
    }
}

func main() {
    http.HandleFunc("/upload", upload)

    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServer ", err)
    }
}
