package db

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql" 
    "fmt"
)

func DBConnect() *sql.DB {
    db, err := sql.Open("mysql", "testing:testing@/bullworth")
    if err != nil {
       fmt.Scanf("Error connecting to DB %v", err) 
    } else {
        fmt.Println("db connected")
    }

    err = db.Ping()
    fmt.Println(err)
    if err != nil {
        fmt.Println("db is not connected")
        fmt.Println(err.Error())
    }
    return db
}
