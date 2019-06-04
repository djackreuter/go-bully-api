package models

import (
    "database/sql"
    "github.com/djackreuter/go-bully-api/db"
    _ "github.com/go-sql-driver/mysql" 
)

type Location struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var conn *sql.DB

func GetLocations() ([]Location, error) {
    conn := db.DBConnect()

    sqlstatement := "SELECT * FROM locations"
    rows, err := conn.Query(sqlstatement)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    result := []Location{}

    for rows.Next() {
        loc := Location{}
        if err := rows.Scan(&loc.ID, &loc.Name); err != nil {
            return nil, err
        }
        result = append(result, loc)
    }
   return result, nil 
}
