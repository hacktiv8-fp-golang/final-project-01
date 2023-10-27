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
// @Tags todos
// @Accept json
// @Produce json
// @Param domain.Todo body domain.TodoRequest true "Todo object to be created"
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
	parsedID, err := strconv.Atoi(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ID must be a number"})
		return
	}

	if err := context.ShouldBindJSON(&todo); err != nil {
		context.JSON(http.StatusUnprocessableEntity, "invalid json body")
		return
	}

	todoResult, err := service.TodoService.UpdateTodo(&todo, parsedID)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, todoResult)
}

func DeleteTodo(context *gin.Context) {
	id := context.Param("id")
	parsedID, err := strconv.Atoi(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ID must be a number"})
		return
	}

	err = service.TodoService.DeleteTodo(parsedID)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Todo with id %d has been successfully deleted", parsedID),
	})
}

func GetAllTodos(context *gin.Context) {

	todoResult, err := service.TodoService.GetAllTodos()

	if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	context.JSON(http.StatusOK, todoResult)
}

func GetTodoByID(context *gin.Context) {

    id := context.Param("id")

    idInt, err := strconv.Atoi(id)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "ID must be a number"})
        return
    }

    todoResult, err := service.TodoService.GetTodoByID(idInt)

    if err != nil {
        context.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Data with ID %+v was not found", id)})
        return
    }

    context.JSON(http.StatusOK, todoResult)
}