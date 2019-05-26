package models

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql" 
    "github.com/djackreuter/go-bully-api/db"
)

type Class struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

type Classes struct {
    Classes []Class `json:"classes"`
}

var con *sql.DB

func GetClasses() Classes {
    con := db.DBConnect()
    sqlstatement := "SELECT id, name FROM classes ORDER BY id"

    rows, err := con.Query(sqlstatement)
    fmt.Println(rows)
    fmt.Println(err)
    if err != nil {
        fmt.Println(err)
    }
    defer rows.Close()

    result := Classes{}

    for rows.Next() {
        class := Class{}
        err2 := rows.Scan(&class.ID, &class.Name)
        if err2 != nil {
            fmt.Println(err2)
        }
        result.Classes = append(result.Classes, class)
    }
    return result
}
