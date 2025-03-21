package handler

import (
	"TestProject/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RegisterUserHandler(r *gin.RouterGroup) {
	r.GET("/all", getAllUsersHandler)
	r.GET("/:id", getUserHandler)
}

// @Summary Get All Users
// @Description 取得所有使用者
// @Tags User
// @Success 200 {array} model.User
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /user/all [get]
func getAllUsersHandler(c *gin.Context) {
	users, err := service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// @Summary Get Single User
// @Description 根據 ID 取得使用者
// @Tags User
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /user/{id} [get]
func getUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	user, err := service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
