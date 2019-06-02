package models

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/djackreuter/go-bully-api/db"
    "strconv"
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

func GetClass(id string) (Class, error) {
    con := db.DBConnect()
    sqlstatement := "SELECT id, name FROM classes WHERE id= ?"
    row := con.QueryRow(sqlstatement, id)

    result := Class{}
    err := row.Scan(&result.ID, &result.Name)
    if err != nil {
        fmt.Println("No Rows", err)
    }
    return result, err
}

func CreateClass(class *Class) (Class, error) {
    con := db.DBConnect()
    var newClass Class
    sqlstatement := "INSERT INTO classes(name) VALUES(?)"
    res, err := con.Exec(sqlstatement, class.Name)
    if err != nil {
        return newClass, err
    }
    id, err := res.LastInsertId()
    if err != nil {
        return newClass, err
    }
    strId := strconv.FormatInt(id, 10)
    newClass, err = GetClass(strId)
    if err != nil {
        return newClass, err
    }
    return newClass, nil
}
