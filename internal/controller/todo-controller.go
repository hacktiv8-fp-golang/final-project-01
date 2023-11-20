package controller

import (
	"github.com/hacktiv8-fp-golang/final-project-01/internal/domain"
	"github.com/hacktiv8-fp-golang/final-project-01/internal/service"
	"github.com/hacktiv8-fp-golang/final-project-01/internal/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateTodo godoc
// @Summary Create a new todo item.
// @Tags Todos
// @Accept json
// @Produce json
// @Param domain.Todo body domain.TodoCreate true "Todo object to be created"
// @Success 201 {object} domain.Todo "Todo created successfully"
// @Failure 400 {object} utils.ErrorResponse "Bad Request"
// @Failure 422 {object} utils.ErrorResponse "Unprocessible Entity"
// @Failure 500 {object} utils.ErrorResponse "Server Error"
// @Router /todos [post]
func CreateTodo(context *gin.Context) {
	var newTodo domain.Todo

	if err := context.ShouldBindJSON(&newTodo); err != nil {
		errHandler := utils.UnprocessibleEntity("Invalid JSON body")
		context.JSON(errHandler.GetStatusCode(), errHandler)
		return
	}

	createdResult, err := service.TodoService.CreateTodo(&newTodo)

	if err != nil {
		context.JSON(err.GetStatusCode(), err)
		return
	}

	context.JSON(http.StatusCreated, createdResult)
}

// UpdateTodo godoc
// @Summary Update a todo item
// @Description Update a todo item by id
// @Tags Todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param domain.Todo body domain.TodoUpdate true "Todo object that needs to be updated"
// @Success 200 {object} domain.Todo
// @Failure 400 {object} utils.ErrorResponse "Bad Request"
// @Failure 404 {object} utils.ErrorResponse "Not Found"
// @Failure 422 {object} utils.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} utils.ErrorResponse "Server Error"
// @Router /todos/{id} [put]
func UpdateTodo(context *gin.Context) {
	id, err := getTodoIdParam(context)

	if err != nil {
		context.JSON(err.GetStatusCode(), err)
		return
	}

	var updatedTodo domain.TodoUpdate

	if err := context.ShouldBindJSON(&updatedTodo); err != nil {
		errHandler := utils.UnprocessibleEntity("Invalid JSON body")
		context.JSON(errHandler.GetStatusCode(), errHandler)
		return
	}

	updatedResult, err := service.TodoService.UpdateTodo(&updatedTodo, id)

	if err != nil {
		context.JSON(err.GetStatusCode(), err)
		return
	}

	context.JSON(http.StatusOK, updatedResult)
}

// DeleteTodo godoc
// @Summary Delete a todo item
// @Description Delete a todo item by ID
// @Tags Todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {string} string "OK"
// @Failure 400 {object} utils.ErrorResponse "Bad Request"
// @Failure 404 {object} utils.ErrorResponse "Not Found"
// @Failure 500 {object} utils.ErrorResponse "Server Error"
// @Router /todos/{id} [delete]
func DeleteTodo(context *gin.Context) {
	id, err := getTodoIdParam(context)

	if err != nil {
		context.JSON(err.GetStatusCode(), err)
		return
	}

	err = service.TodoService.DeleteTodo(id)

	if err != nil {
		context.JSON(err.GetStatusCode(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Todo with id %d has been successfully deleted", id),
	})
}

// GetAllTodos godoc
// @Summary Get all todos
// @Description Retrieve a list of all todos
// @Tags Todos
// @Accept json
// @Produce json
// @Success 200 {object} []domain.Todo
// @Failure 400 {object} utils.ErrorResponse "Bad Request"
// @Failure 404 {object} utils.ErrorResponse "Not Found"
// @Failure 500 {object} utils.ErrorResponse "Server Error"
// @Router /todos [get]
func GetAllTodos(context *gin.Context) {
	todos, err := service.TodoService.GetAllTodos()

	if err != nil {
		context.JSON(err.GetStatusCode(), err)
		return
	}

	context.JSON(http.StatusOK, todos)
}

// GetTodoByID godoc
// @Summary Get a todo item
// @Description Get a todo item by ID
// @Tags Todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} domain.Todo
// @Failure 400 {object} utils.ErrorResponse "Bad Request"
// @Failure 404 {object} utils.ErrorResponse "Not Found"
// @Router /todos/{id} [get]
func GetTodoByID(context *gin.Context) {
	id, err := getTodoIdParam(context)

	if err != nil {
		context.JSON(err.GetStatusCode(), err)
		return
	}

	todo, err := service.TodoService.GetTodoByID(id)

	if err != nil {
		context.JSON(err.GetStatusCode(), err)
		return
	}

	context.JSON(http.StatusOK, todo)
}

func getTodoIdParam(context *gin.Context) (int, utils.Error) {
	idParam := context.Param("id")
	todoId, err := strconv.Atoi(idParam)

	if err != nil {
		return int(0), utils.BadRequest("ID param must be an integer")
	}

	return int(todoId), nil
}
