package db

import (
	"TestProject/config/logger"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		logger.Fatal("db init error", zap.Error(err))
	} else {
		logger.Info("db init success")
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, age INTEGER)`)
	if err != nil {
		logger.Fatal("db init error", zap.Error(err))
	}

	_, _ = db.Exec(`INSERT INTO users (name, age) VALUES ('Hiro', 22), ('Zoe', 21), ('Ken', 99)`)
}

func GetDB() *sql.DB {
	return db
}
