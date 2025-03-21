package db

import (
	d "TestProject/config/db"
	"TestProject/config/logger"
	"TestProject/model"
	"database/sql"
)

var db *sql.DB

func GetUserByID(id int) (*model.User, error) {
	db = d.GetDB()

	row := db.QueryRow("SELECT id, name, age FROM users WHERE id = ?", id)
	var u model.User
	err := row.Scan(&u.ID, &u.Name, &u.Age)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func GetAllUsers() ([]model.User, error) {
	db = d.GetDB()

	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			logger.Fatal("rows close error")
		}
	}(rows)

	var users []model.User
	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.ID, &u.Name, &u.Age)
		if err != nil {
			continue
		}
		users = append(users, u)
	}
	return users, nil
}
