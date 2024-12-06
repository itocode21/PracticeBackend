package body

import (
	"todoapi/mid"

	"github.com/gin-gonic/gin"
)

func RegisterServer(server gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(mid.Authentication)
	authenticated.POST("/todos", createTodos)
	authenticated.PUT("/todos/:id", updateTodos)
	authenticated.DELETE("/todos/:id", deleteTodos)
	authenticated.GET("/todos", getTodos)

	server.POST("/register", createUser)
	server.POST("/login", login)
}
