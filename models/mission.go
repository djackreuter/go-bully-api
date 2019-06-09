package models

import (
    _ "database/sql"
    "strconv"
    "github.com/djackreuter/go-bully-api/db"
	_ "github.com/go-sql-driver/mysql"
)

type Mission struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}


func GetMissions() ([]*Mission, error) {
    conn := db.DBConnect()
    missions := []*Mission{}

    sqlstatement := "SELECT * FROM missions;"

    rows, err := conn.Query(sqlstatement)
    if err != nil {
        return missions, err
    }

    defer rows.Close()

    for rows.Next() {
        mission := &Mission{}
        if err := rows.Scan(&mission.ID, &mission.Name); err != nil {
            return missions, err
        }
        missions = append(missions, mission)
    }
    return missions, nil
}

func GetMission(id string) (*Mission, error) {
    conn := db.DBConnect() 
    mission := &Mission{}
    sqlstatement := "SELECT * FROM missions WHERE id = ?;"

    row := conn.QueryRow(sqlstatement, id)
    if err := row.Scan(&mission.ID, &mission.Name); err != nil {
        return mission, err
    }
    return mission, nil
}

func CreateMission(mission *Mission) (*Mission, error) {
    conn := db.DBConnect()
    var newMission *Mission
    sqlstatement := "INSERT INTO missions(name) VALUES(?);"

    res, err := conn.Exec(sqlstatement, mission.Name)
    if err != nil {
        return newMission, err
    }
    id, err := res.LastInsertId()
    if err != nil {
        return newMission, err
    }
    return GetMission(strconv.FormatInt(id, 10))
}

func UpdateMission(mission *Mission) (*Mission, error) {
    conn := db.DBConnect()
    var newMission *Mission
    sqlstatement := "UPDATE missions SET name = ? WHERE id = ?;"

    _, err := conn.Exec(sqlstatement, mission.Name, mission.ID)
    if err != nil {
        return newMission, err
    }
    return GetMission(strconv.Itoa(mission.ID))
}

func DeleteMission(mission *Mission) error {
    conn := db.DBConnect()
    sqlstatement := "DELETE FROM missions WHERE id = ?;"

    _, err := conn.Exec(sqlstatement, mission.ID)
    if err != nil {
        return err
    }
    return nil
}
