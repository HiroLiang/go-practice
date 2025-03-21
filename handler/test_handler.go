package handler

import (
	"TestProject/config/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func RegisterTestHandler(r *gin.RouterGroup) {
	r.GET("/ping", pingHandler)
	r.GET("/hello/:name", helloHandler)
}

// @Summary Ping
// @Description Health check
// @Tags Test
// @Success 200 {object} map[string]string
// @Router /test/ping [get]
func pingHandler(c *gin.Context) {
	logger.Info("Ping called")
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// @Summary Hello
// @Description 回應 Hello Name
// @Tags Test
// @Param name path string true "使用者名稱"
// @Success 200 {object} map[string]string
// @Router /test/hello/{name} [get]
func helloHandler(c *gin.Context) {
	name := c.Param("name")
	logger.Info("Hello called", zap.String("name", name))
	c.JSON(http.StatusOK, gin.H{"message": "Hello " + name})
}
