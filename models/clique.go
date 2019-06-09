package models

import (
    _ "database/sql"
    "strconv"
    "github.com/djackreuter/go-bully-api/db"
	_ "github.com/go-sql-driver/mysql"
)

type Clique struct {
    ID       int        `json:"id"`
    Name     string     `json:"name"`
}

func GetCliques() ([]*Clique, error) {
    conn := db.DBConnect()
    var cliques []*Clique
    sqlstatement := "SELECT * FROM cliques;"

    rows, err := conn.Query(sqlstatement)
    if err != nil {
        return cliques, err
    }
    defer rows.Close()
    
    for rows.Next() {
        clique := &Clique{}
        if err := rows.Scan(&clique.ID, &clique.Name); err != nil {
            return cliques, err
        }
        cliques = append(cliques, clique)
    }
    return cliques, nil
}

func GetClique(id string) (*Clique, error) {
    conn := db.DBConnect() 
    clique := &Clique{}
    sqlstatement := "SELECT * FROM cliques WHERE id = ?;"

    row := conn.QueryRow(sqlstatement, id)
    if err := row.Scan(&clique.ID, &clique.Name); err != nil {
        return clique, err
    }
    return clique, nil
}   

func CreateClique(clique *Clique) (*Clique, error) {
    conn := db.DBConnect()
    var newClique *Clique
    sqlstatement := "INSERT INTO cliques(name) VALUES(?);"

    res, err := conn.Exec(sqlstatement, clique.Name)
    if err != nil {
        return newClique, err
    }
    id, err := res.LastInsertId()
    if err != nil {
        return newClique, err
    }
    return GetClique(strconv.FormatInt(id, 10))
}

func UpdateClique(clique *Clique) (*Clique, error) {
    conn := db.DBConnect()
    var newClique *Clique
    sqlstatement := "UPDATE cliques SET name = ? WHERE id = ?;"

    _, err := conn.Exec(sqlstatement, clique.Name, clique.ID)
    if err != nil {
        return newClique, err
    }
    return GetClique(strconv.Itoa(clique.ID))
}

func DeleteClique(clique *Clique) error {
    conn := db.DBConnect()
    sqlstatement := "DELETE FROM cliques WHERE id = ?;"

    _, err := conn.Exec(sqlstatement, clique.ID)
    if err != nil {
        return err
    }
    return nil
}
