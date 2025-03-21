package service

import (
	"TestProject/db"
	"TestProject/model"
)

func GetUserByID(id int) (*model.User, error) {
	return db.GetUserByID(id)
}

func GetAllUsers() ([]model.User, error) {
	return db.GetAllUsers()
}
