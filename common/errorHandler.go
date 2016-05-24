package common

import (
	"github.com/gin-gonic/gin"

)

type (
	appError struct {
		Error string `json:"error"`
		Message string `json:"message"`
		HttpStatus int `json:"status"`
	}

	errorResource struct {
		Data appError `json:"data"`
		Success bool `json:"success"`
	}
)

func RenderError(c *gin.Context, errCode int, handlerError error, message string){
	errorResponse := errorResource{
		Success: false,
		Data: appError{
			Error: handlerError.Error(),
			Message: message,
			HttpStatus: errCode,
		},
	}
	//log.Panic(handlerError)

	c.JSON(errCode, errorResponse)
}