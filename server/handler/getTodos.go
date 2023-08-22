package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiburciohugo/go-react-todo/schemas"
)

func GetTodosHandler(c *gin.Context) {
	todos := []schemas.Todo{}

	if err := db.Find(&todos).Error; err != nil {
		logger.Errorf("error fetching todos: %v", err.Error())
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, http.StatusOK, todos)
}
