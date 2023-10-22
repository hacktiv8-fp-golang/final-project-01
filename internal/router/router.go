package router

import "github.com/gin-gonic/gin"

var PORT = ":8080"

func StartServer() {
	router := gin.Default()

	todosRouter := router.Group("/todos")
	{
		todosRouter.POST("/")
		todosRouter.GET("/")
		todosRouter.GET("/:id")
		todosRouter.PUT("/:id")
		todosRouter.DELETE("/:id")
	}

	router.Run(PORT)
}