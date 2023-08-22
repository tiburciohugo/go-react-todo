package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiburciohugo/go-react-todo/schemas"
)

func DeleteTodoHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	todo := schemas.Todo{}

	if err := db.First(&todo, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("todo with id: %s not found", id))
		return
	}

	if err := db.Delete(&todo).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("error deleting todo with id: %s", id))
		return
	}
	sendSuccess(c, http.StatusAccepted, todo)
}
