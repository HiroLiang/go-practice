// @title        GoLang API Example
// @version      1.0
// @description  這是用 gin 製作的範例 API 文件
// @host         localhost:8080
// @BasePath     /
package main

import (
	_ "TestProject/docs"
	"TestProject/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic("config read error: " + err.Error())
	}
}

// PingHandler @Summary      Ping
// @Description  健康檢查 API
// @Tags         Health
// @Success      200 {object} map[string]string
// @Router       /ping [get]
func PingHandler(c *gin.Context) {
	logger.Info("Ping called")
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// HelloHandler @Summary      Hello
// @Description  回應 Hello Name
// @Tags         Hello
// @Param        name path string true "使用者名稱"
// @Success      200 {object} map[string]string
// @Router       /hello/{name} [get]
func HelloHandler(c *gin.Context) {
	name := c.Param("name")
	logger.Info("Hello called", zap.String("name", name))
	c.JSON(http.StatusOK, gin.H{"message": "Hello " + name})
}

func main() {
	util.Hello()

	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	initConfig()

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", PingHandler)
	r.GET("/hello/:name", HelloHandler)

	r.Run(":" + viper.GetString("port"))
}
