// C:\GoProject\src\eShop\pkg\controllers\users.go

package controllers

import (
	"eShop/errs"
	"eShop/logger"
	"eShop/models"
	"eShop/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	logger.Info.Printf("IP: [%s] requested list of all users\n", c.ClientIP()) // Логируем IP при запросе всех пользователей...

	users, err := service.GetAllUsers()
	if err != nil {
		logger.Error.Printf("[controllers.GetAllUsers] error getting all users: %v\n", err)
		handleError(c, err)
		return
	}

	logger.Info.Printf("IP: [%s] got list of all users\n", c.ClientIP()) // Логируем IP при успешной выдаче списка...
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetUserByID] invalid user_id path parameter: %s, IP: [%s]\n", c.Param("id"), c.ClientIP())
		handleError(c, errs.ErrValidationFailed)
		return
	}

	logger.Info.Printf("IP: [%s] requested user with ID: %d\n", c.ClientIP(), id) // Логируем IP и запрашиваемый ID пользователя

	user, err := service.GetUserByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("IP: [%s] got user with ID: %d\n", c.ClientIP(), id) // Логируем IP и успешное получение пользователя
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err := service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully!!!",
	})
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	user.ID = uint(id)

	c.JSON(http.StatusOK, user)
}
