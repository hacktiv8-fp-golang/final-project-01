package controller

import (
	"final-project-01/internal/database"
	"final-project-01/internal/domain"
	"final-project-01/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTodo(context *gin.Context) {
	var todo domain.Todo

	if err := context.ShouldBindJSON(&todo); err != nil {
		context.JSON(http.StatusUnprocessableEntity, "invalid json body")
		return
	}

	todoResult, err := service.TodoService.CreateTodo(&todo)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusCreated, todoResult)
}

func GetAllData(context *gin.Context) {
	var data []domain.Todo
	db := database.GetDB()

	if err := db.Find(&data).Error; err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}
	context.JSON(http.StatusOK, data)
}

func GetDataByID(context *gin.Context) {
	var data domain.Todo
	db := database.GetDB()
	id := context.Param("id")

	if err := db.First(&data, id).Error; err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}
	context.JSON(http.StatusOK, data)
}