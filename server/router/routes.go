package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	basePath := "/api/v1"
	v1 := r.Group(basePath)
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		v1.GET("/todos", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "GET todos",
			})
		})
		v1.GET("/todo/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "GET todo",
			})
		})
		v1.POST("/todo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "POST todo",
			})
		})
		v1.PUT("/todo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "PUT todo",
			})
		})
		v1.DELETE("/todo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "DELETE todo",
			})
		})
	}
}
