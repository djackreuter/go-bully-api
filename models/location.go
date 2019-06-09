package models

import (
	"database/sql"
    "strconv"

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

	sqlstatement := "SELECT * FROM locations;"
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

func GetLocation(id string) (*Location, error) {
	conn := db.DBConnect()

	sqlstatement := "SELECT * FROM locations WHERE id = ?;"
	row := conn.QueryRow(sqlstatement, id)
	loc := &Location{}

	err := row.Scan(&loc.ID, &loc.Name)
	if err != nil {
		return loc, err
	}
	return loc, nil
}

func AddLocation(loc *Location) (*Location, error) {
	conn := db.DBConnect()
    var newLoc *Location

    sqlstatement := "INSERT INTO locations (name) VALUES(?);"

    res, err := conn.Exec(sqlstatement, loc.Name)
    if err != nil {
        return newLoc, err
    }
    intId, err := res.LastInsertId()
    if err != nil {
        return newLoc, err
    }
    return GetLocation(strconv.FormatInt(intId, 16))
}

func UpdateLocation(loc *Location) (*Location, error) {
    conn := db.DBConnect()
    var updatedLoc *Location

    sqlstatement := "UPDATE locations SET name = ? WHERE id = ?;"
    _, err := conn.Exec(sqlstatement, loc.Name, loc.ID)
    if err != nil {
        return updatedLoc, err
    }
    return GetLocation(strconv.Itoa(loc.ID))
}

func DeleteLocation(loc *Location) error {
    conn := db.DBConnect()
    sqlstatement := "DELETE FROM locations WHERE id = ?;"

    _, err := conn.Exec(sqlstatement, loc.ID)
    if err != nil {
        return err
    }
    return nil
}
