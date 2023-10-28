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

// UpdateTodo godoc
// @Summary Update a todo item
// @Description Update a todo item by ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "todo ID"
// @Param domain.Todo body domain.TodoRequest true "Todo object that needs to be updated"
// @Success 200 {object} domain.Todo
// @Failure 400 {string} string "Bad Request"
// @Failure 422 {string} string "Unprocessable Entity"
// @Router /todos/{id} [put]
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

// DeleteTodo godoc
// @Summary Delete a todo item
// @Description Delete a todo item by ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Router /todos/{id} [delete]
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

// GetAllTodos godoc
// @Summary Get all todos
// @Description Retrieve a list of all todos
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {object} []domain.Todo
// @Failure 400 {string} string "Bad Request"
// @Router /todos [get]
func GetAllTodos(context *gin.Context) {

	todoResult, err := service.TodoService.GetAllTodos()

	if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	context.JSON(http.StatusOK, todoResult)
}

// GetTodoByID godoc
// @Summary Get a todo item
// @Description Get a todo item by ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} domain.Todo
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Router /todos/{id} [get]
func GetTodoByID(context *gin.Context) {

    id := context.Param("id")

    parsedID, err := strconv.Atoi(id)

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "ID must be a number"})
        return
    }

    todoResult, err := service.TodoService.GetTodoByID(parsedID)

    if err != nil {
        context.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Data with ID %+v was not found", id)})
        return
    }

    context.JSON(http.StatusOK, todoResult)
}