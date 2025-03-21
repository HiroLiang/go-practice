// @title        GoLang API Practice
// @version      1.0.0
// @description  練習用 Go lang 測試
// @host         localhost:8080
// @BasePath     /
package main

import (
	"TestProject/config/api"
	"TestProject/config/db"
	_ "TestProject/config/db"
	"TestProject/config/logger"
	_ "TestProject/docs"
	"github.com/joho/godotenv"
)

// 初始化環境變數 (.env 要自己加)
func init() {
	_ = godotenv.Load() // 讀取 .env（建議放 init）
}

func main() {
	// 初始化全域 Logger
	logger.Init()
	defer logger.Stop()

	db.InitDB()

	// 啟動 Http Server
	api.StartServer()
}
