package router

import (
	"TestProject/handler"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initSwagConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic("config read error: " + err.Error())
	}
}

func SetupRouter() *gin.Engine {
	// 初始化 Swagger 設定
	initSwagConfig()

	// 使用預設 Rest 並註冊群組
	r := gin.Default()

	// 註冊 Swagger
	swagger := r.Group("/swagger")
	swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 註冊 Test APIs
	test := r.Group("/test")
	handler.RegisterTestHandler(test)

	// 註冊 User 相關 APIs
	user := r.Group("/user")
	handler.RegisterUserHandler(user)

	return r
}
