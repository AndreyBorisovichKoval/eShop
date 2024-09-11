// C:\GoProject\src\eShop\pkg\controllers\routes.go

package controllers

import (
	"eShop/configs"
	"eShop/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		userG.POST("", SignUp)
		userG.GET("", GetAllUsers)
		userG.GET("/:id", GetUserByID)
		userG.PUT("/:id", UpdateUserByID)
		userG.DELETE("/:id", SoftDeleteUserByID)
		userG.PUT("/:id/restore", RestoreUserByID)

		// sellerG.POST("", CreateSellers)
		// sellerG.PATCH("/:id", PatchSellerByID)
		// sellerG.DELETE("/harddelete/:id", HardDeleteNoteByID)
		// sellerG.DELETE("/softdelete/:id", SoftDeleteNoteByID)
		// sellerG.PUT("/restore/:id", RestoreNoteByID)
	}

	return router
}
