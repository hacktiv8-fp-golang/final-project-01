package router

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/hacktiv8-fp-golang/final-project-01/docs"
	"github.com/hacktiv8-fp-golang/final-project-01/internal/controller"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartServer() {
	router := gin.Default()

	todosRouter := router.Group("/todos")
	{
		todosRouter.POST("/", controller.CreateTodo)
		todosRouter.GET("/", controller.GetAllTodos)
		todosRouter.GET("/:id", controller.GetTodoByID)
		todosRouter.PUT("/:id", controller.UpdateTodo)
		todosRouter.DELETE("/:id", controller.DeleteTodo)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	var PORT = os.Getenv("PORT")

	router.Run(":" +PORT)
}
