package router

import (
	"final-project-01/internal/controller"

	"github.com/gin-gonic/gin"
)

var PORT = ":8080"

func StartServer() {
	router := gin.Default()

	todosRouter := router.Group("/todos")
	{
		todosRouter.POST("/", controller.CreateTodo)
		todosRouter.GET("/")
		todosRouter.GET("/:id")
		todosRouter.PUT("/:id", controller.UpdateTodo)
		todosRouter.DELETE("/:id", controller.DeleteTodo)
	}

	router.Run(PORT)
}
