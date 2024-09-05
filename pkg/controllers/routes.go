// C:\GoProject\src\eShop\pkg\controllers\routes.go

package controllers

import (
	"eShop/configs"
	"eShop/logger"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingPong(c *gin.Context) {
	logger.Info.Println("PingPong route called...")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func RunRoutes() error {
	router := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)

	router.GET("/ping", PingPong)

	// authG := router.Group("/auth")
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}

	sellerG := router.Group("/sellers")
	{
		sellerG.GET("", GetAllSellers)
		// sellerG.GET("/:id", GetNoteByID)
		// sellerG.POST("", CreateNote)
		// sellerG.PUT("/:id", UpdateNoteByID)
		// sellerG.PATCH("/:id", PatchNoteByID)
		// sellerG.DELETE("/harddelete/:id", HardDeleteNoteByID)
		// sellerG.DELETE("/softdelete/:id", SoftDeleteNoteByID)
		// sellerG.PUT("/restore/:id", RestoreNoteByID)
	}

	err := router.Run(fmt.Sprintf("%s:%s", configs.AppSettings.AppParams.ServerURL, configs.AppSettings.AppParams.PortRun))
	if err != nil {
		logger.Error.Printf("Server failed to start: %v", err)
		return err
	}

	logger.Info.Printf("Server started on port: %s...\n", configs.AppSettings.AppParams.PortRun)
	return nil
}
