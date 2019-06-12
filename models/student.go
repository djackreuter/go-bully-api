package models

import (
	"fmt"
	"strconv"

	"github.com/djackreuter/go-bully-api/db"
	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	CliqueID int    `json:"cliqueId"`
}

func intToString(num int) string {
	return strconv.Itoa(num)
}

func GetAllStudents() ([]*Student, error) {
	conn := db.DBConnect()
	sqlstatement := "SELECT * FROM students;"
	var students []*Student

	rows, err := conn.Query(sqlstatement)
	if err != nil {
		return students, err
	}
	defer rows.Close()

	for rows.Next() {
		student := &Student{}
		if err := rows.Scan(&student.ID, &student.Name, &student.CliqueID); err != nil {
			return students, err
		}
		students = append(students, student)
	}
	return students, nil
}

func GetCliqueStudents(cliqueId string) (interface{}, error) {
	conn := db.DBConnect()
	sqlstatement := "SELECT c.name, s.id, s.name, s.clique_id FROM students s JOIN cliques c ON s.clique_id = c.id AND clique_id = ?;"
	var cliqueStudents []interface{}

	rows, err := conn.Query(sqlstatement, cliqueId)
	if err != nil {
		return cliqueStudents, err
	}
	defer rows.Close()
	for rows.Next() {
		students := &Student{}
		clique := &Clique{}
		if err := rows.Scan(&clique.Name, &students.ID, &students.Name, &students.CliqueID); err != nil {
			return cliqueStudents, err
		}
		cliqueStudents = append(cliqueStudents, students, clique)
	}
	return cliqueStudents, nil
}

func GetCliqueStudent(clique_id, id string) (*Student, error) {
	conn := db.DBConnect()

	sqlstatement := "SELECT * FROM students WHERE clique_id = ? AND id = ?;"
	row := conn.QueryRow(sqlstatement, clique_id, id)
	student := &Student{}

	err := row.Scan(&student.ID, &student.Name, &student.CliqueID)
	if err != nil {
		return student, err
	}
	return student, nil
}

func CreateCliqueStudent(student *Student) (*Student, error) {
	conn := db.DBConnect()

	sqlstatement := "INSERT INTO students (name, clique_id) VALUES(?, ?);"
	res, err := conn.Exec(sqlstatement, student.Name, student.CliqueID)
	i64Id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("Err here")
		return nil, err
	}

	return GetCliqueStudent(intToString(student.CliqueID), intToString(int(i64Id)))
}

func UpdateCliqueStudent(student *Student) (*Student, error) {
	conn := db.DBConnect()

	sqlstatement := "UPDATE students SET name = ?, clique_id = ? WHERE id = ?;"
	_, err := conn.Exec(sqlstatement, student.Name, student.CliqueID, student.ID)
	if err != nil {
		return nil, err
	}
	s, err := GetCliqueStudent(intToString(student.CliqueID), intToString(student.ID))
	if err != nil {
		return s, err
	}
	return s, nil
}

func DeleteCliqueStudent(id string) error {
	conn := db.DBConnect()
	sqlstatement := "DELETE FROM students WHERE id = ?;"

	_, err := conn.Exec(sqlstatement, id)
	if err != nil {
		return err
	}
	return nil
}
