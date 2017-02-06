package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func checkError(err error) {
    if err != nil {
        panic(err
    }
}

func main() {
    db, err := sql.Open("sqlite3", "./test.db")
    checkError(err)

    // insert item
    stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?, ?, ?)")
    checkError(err)

    res, err := stmt.Exec("zzh", "dev", "2017")
    checkError(err)

    id, err := res.LastInsertId()
    checkError(err)

    fmt.Println(id)
    // update
    stmt, err := db.Prepare("update userinfo set username=? where uid=?")
    checkError(err)
    res, err := stmt.Exec("zzhupdate", id)
    checkError(err)

    affect, err := res.RowsAffected()
    checkError(err)

    fmt.Println(affect)

    // select db
    rows, err := db.Query("SELECT * FROM userinfo")
    checkError(err)

    for rows.Next() {
        var uid int
        var username string
        var department string
        var created string
        err = rows.Scan(&uid, &username, &department, &created)
        checkError(err)
        fmt.Println(uid)
        fmt.Println(username)
        fmt.Println(department)
        fmt.Println(created)
    }
}
