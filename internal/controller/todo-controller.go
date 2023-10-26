package controller

import (
	"final-project-01/internal/domain"
	"final-project-01/internal/service"
	"net/http"
	"strconv"

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

func UpdateTodo(context *gin.Context) {
	var todo domain.Todo
	id := context.Param("id")

	if err := context.ShouldBindJSON(&todo); err != nil {
		context.JSON(http.StatusUnprocessableEntity, "invalid json body")
		return
	}

	todoResult, err := service.TodoService.UpdateTodo(&todo, id)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, todoResult)
}

func DeleteTodo(context *gin.Context) {
	id := context.Param("id")

	todoResult, err := service.TodoService.DeleteTodo(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, todoResult)
}

func GetAllData(context *gin.Context) {

	todoResult, err := service.TodoService.GetAllData()

	if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	context.JSON(http.StatusOK, todoResult)
}

func GetDataByID(context *gin.Context) {

    id := context.Param("id")

    idInt, err := strconv.Atoi(id)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
        return
    }

    todoResult, err := service.TodoService.GetDataByID(idInt)

    if err != nil {
        context.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
        return
    }

    context.JSON(http.StatusOK, todoResult)
}