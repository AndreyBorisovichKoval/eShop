// C:\GoProject\src\eShop\pkg\controllers\helpers.go

package controllers

import (
	"coinkeeper/errs"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleError(c *gin.Context, err error) {
	if errors.Is(err, errs.ErrUsernameUniquenessFailed) ||
		errors.Is(err, errs.ErrIncorrectUsernameOrPassword) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else if errors.Is(err, errs.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else if errors.Is(err, errs.ErrPermissionDenied) {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	} else {

		c.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrSomethingWentWrong.Error()})
	}
}
