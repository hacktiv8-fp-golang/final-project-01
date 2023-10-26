package router

import (
	"final-project-01/internal/controller"
	_ "final-project-01/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/files"
)

var PORT = ":8080"

// @Title Todos API
// @version 1.0
// @description This API service allows you to manage todo items.
// @host localhost:8080
// @BasePath /
func StartServer() {
	router := gin.Default()

	todosRouter := router.Group("/todos")
	{
		todosRouter.POST("/", controller.CreateTodo)
		todosRouter.GET("/")
		todosRouter.GET("/:id")
		todosRouter.PUT("/:id")
		todosRouter.DELETE("/:id")
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(PORT)
}