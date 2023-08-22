package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiburciohugo/go-react-todo/handler"
)

func initializeRoutes(r *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"
	v1 := r.Group(basePath)
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		v1.GET("/todos", handler.GetTodosHandler)
		v1.GET("/todo/", handler.GetTodoByIdHandler)
		v1.POST("/todo", handler.CreateTodoHandler)
		v1.PUT("/todo", handler.UpdateTodoHandler)
		v1.DELETE("/todo", handler.DeleteTodoHandler)
	}
}
