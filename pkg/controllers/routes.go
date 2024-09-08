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
	{
		userG.POST("/create", SignUp)
		userG.GET("", GetAllUsers)
		// sellerG.GET("/:id", GetSellerByID)
		// sellerG.POST("", CreateSellers)
		// sellerG.PUT("/:id", UpdateSellerByID)
		// sellerG.PATCH("/:id", PatchSellerByID)
		// sellerG.DELETE("/harddelete/:id", HardDeleteNoteByID)
		// sellerG.DELETE("/softdelete/:id", SoftDeleteNoteByID)
		// sellerG.PUT("/restore/:id", RestoreNoteByID)
	}

	// err := router.Run(fmt.Sprintf("%s:%s", configs.AppSettings.AppParams.ServerURL, configs.AppSettings.AppParams.PortRun))
	// if err != nil {
	// 	logger.Error.Printf("Server failed to start: %v", err)
	// 	return err
	// }

	// logger.Info.Printf("Server started on port: %s...\n", configs.AppSettings.AppParams.PortRun)
	// return nil

	return router
}
