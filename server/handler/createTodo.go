package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiburciohugo/go-react-todo/schemas"
)

func CreateTodoHandler(c *gin.Context) {
	request := CreateTodoRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("error validating request: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	todo := schemas.Todo{
		Title:       request.Title,
		Description: request.Description,
		Completed:   request.Completed,
	}

	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("error creating todo: %v", err.Error())
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, http.StatusCreated, todo)
}
