// C:\GoProject\src\eShop\pkg\controllers\routes.go

package controllers

import (
	"eShop/configs"
	_ "eShop/docs"
	"eShop/logger"

	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// localhost:8585/swagger/index.html

// PingPong handles the ping request and responds with a pong message...
func PingPong(c *gin.Context) {
	logger.Info.Printf("Route '%s' called with method '%s'", c.FullPath(), c.Request.Method)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func InitRoutes() *gin.Engine {
	router := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping", PingPong)

	authG := router.Group("/auth")
	{
		// authG.POST("/sign-up", SignUp)
		authG.POST("/sign-in", SignIn)
	}

	userG := router.Group("/users", checkUserAuthentication)
	// userG := router.Group("/users")
	{
		// userG.POST("/create", SignUp)
		// userG.POST("/create", SignUp, checkUserAuthentication)
		userG.POST("", CreateUser)                            // Регистрация нового пользователя...
		userG.GET("", GetAllUsers)                        // Получение списка всех активных пользователей...
		userG.GET("/deleted", GetAllDeletedUsers)         // Получение списка всех удалённых пользователей...
		userG.GET("/:id", GetUserByID)                    // Получение данных пользователя по его ID...
		userG.PATCH("/:id", UpdateUserByID)               // Обновление данных пользователя по его ID...
		userG.PATCH("/:id/block", BlockUserByID)          // Блокировка пользователя по его ID...
		userG.PATCH("/:id/unblock", UnblockUserByID)      // Разблокировка пользователя по его ID...
		userG.PATCH("/:id/restore", RestoreUserByID)      // Восстановление удалённого пользователя...
		userG.DELETE("/:id/soft", SoftDeleteUserByID)     // Софт удаление пользователя по его ID...
		userG.DELETE("/:id/hard", HardDeleteUserByID)     // Полное удаление пользователя по его ID (хард удаление)...
		userG.PATCH("/password", ChangeOwnPassword)       // Обновление пароля текущего пользователя...
		userG.PATCH("/:id/reset-password", ResetPassword) // Сброс пароля пользователя (доступно только администратору)...
	}

	return router
}
