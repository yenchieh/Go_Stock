package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func renderJSONError(c *gin.Context, err error, responseCode int, message string) {
	logrus.Errorf("%+v", err)

	c.JSON(responseCode, gin.H{
		"error":   err.Error(),
		"message": message,
	})
	return
}
