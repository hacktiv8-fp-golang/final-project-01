package controller

import (
	"final-project-01/internal/domain"
	"final-project-01/internal/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateTodo godoc
// @Summary Create a new todo item.
// @Accept json
// @Produce json
// @Param domain.Todo body domain.Todo true "Todo object to be created"
// @Success 201 {object} domain.Todo "Todo created successfully"
// @Failure 400 {string} string "Bad Request: invalid request data"
// @Failure 422 {string} string "Unprocessable Entity: "invalid json body"
// @Router /todos [post]
func CreateTodo(context *gin.Context) {
	var todo domain.Todo

	if err := context.ShouldBindJSON(&todo); err != nil {
		context.JSON(http.StatusUnprocessableEntity, "Invalid JSON body")
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

	err := service.TodoService.DeleteTodo(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Todo with id %+v has been successfully deleted", id),
	})
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