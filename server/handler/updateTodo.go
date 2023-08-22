package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiburciohugo/go-react-todo/schemas"
)

func UpdateTodoHandler(c *gin.Context) {
	request := UpdateTodoRequest{}

	c.BindJSON(request)

	if err := request.Validate(); err != nil {
		logger.Errorf("error validating request: %s", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

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

	if request.Title != "" {
		todo.Title = request.Title
	}
	if request.Description != "" {
		todo.Description = request.Description
	}
	if request.Completed {
		todo.Completed = request.Completed
	}

	if err := db.Save(&todo).Error; err != nil {
		logger.Errorf("error updating todo: %s", err.Error())
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("error updating todo: %s", err.Error()))
		return
	}
	sendSuccess(c, http.StatusAccepted, todo)
}
