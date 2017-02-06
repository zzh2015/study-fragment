package main

import (
    "fmt"
    "html/template"
    "net/http"
    "strings"
    "log"
)

// http://localhost:9090
func sayHelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()

    fmt.Println(r.Form) // request info
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])

    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }

    fmt.Fprintf(w, "Hello go web!") // response
}

// http://localhost:9090/login
func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        log.Println(t.Execute(w, nil))
    } else {
        r.ParseForm()
        fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
        fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
        // fmt.Println("username:", r.Form["username"])
        // fmt.Println("password:", r.Form["password"])
        // reply tips
        var tips string = "welcom access my website,"
        tips += strings.Join(r.Form["username"],"")
        // fmt.Fprintf(w, tips)
        template.HTMLEscape(w, []byte(r.Form.Get("username")))
    }
    // 获取表单内容
    /*
        getInt, err := strconv.Atoi(r.Form.Get("age")
        if err != nil {
            //
        }
        if getInt > 100 {
            //
        }
    */
    // 正则表达形式,略慢
    /*
        if m, _ := regexp.MatchString("^[0-9]+$", r.From.Get("Age")); !m {
        }
    */
}

func main() {
    http.HandleFunc("/", sayHelloName) // 访问的路由
    http.HandleFunc("/login", login)
    err := http.ListenAndServe(":9090", nil)

    if err != nil {
        log.Fatal("ListenAndServer ", err)
    }
}
