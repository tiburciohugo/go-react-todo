package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	schemas "github.com/tiburciohugo/go-react-todo/schemas"
)

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

type SuccessResponse struct {
	Message string               `json:"message"`
	Data    schemas.TodoResponse `json:"data"`
}

func sendError(c *gin.Context, code int, msg string) {
	c.Header("Content-Type", "application/json")
	c.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(c *gin.Context, op int, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %d successful", op),
		"data":    data,
	})
}
