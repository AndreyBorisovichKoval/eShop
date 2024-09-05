// C:\GoProject\src\eShop\pkg\controllers\routes.go

package controllers

import (
	"eShop/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingPong(c *gin.Context) {
	logger.Info.Println("PingPong route called")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func RunRoutes() error {
	router := gin.Default()
	//gin.SetMode(gin.ReleaseMode)

	router.GET("/ping", PingPong)

	noteG := router.Group("/notes")
	{
		noteG.GET("", GetAllNotes)
		noteG.GET("/:id", GetNoteByID)
		noteG.POST("", CreateNote)
		noteG.PUT("/:id", UpdateNoteByID)
		noteG.PATCH("/:id", PatchNoteByID)
		noteG.DELETE("/harddelete/:id", HardDeleteNoteByID)
		noteG.DELETE("/softdelete/:id", SoftDeleteNoteByID)
		noteG.PUT("/restore/:id", RestoreNoteByID)
	}

	err := router.Run(":8585")
	if err != nil {
		logger.Error.Printf("Server failed to start: %v", err)
		return err
	}
	logger.Info.Println("Server started on port 8585.")
	return nil
}
