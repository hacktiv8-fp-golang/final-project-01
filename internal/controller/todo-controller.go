package controller

import (
	"final-project-01/internal/domain"
	"final-project-01/internal/service"
	"net/http"

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