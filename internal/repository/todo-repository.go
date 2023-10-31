package repository

import (
	"github.com/hacktiv8-fp-golang/final-project-01/internal/database"
	"github.com/hacktiv8-fp-golang/final-project-01/internal/domain"
	"github.com/hacktiv8-fp-golang/final-project-01/internal/utils"
	"fmt"
)

type todoDomainRepo interface {
	CreateTodo(*domain.Todo) (*domain.Todo, utils.Error)
	UpdateTodo(*domain.Todo, int) (*domain.Todo, utils.Error)
	DeleteTodo(int) (utils.Error)
	GetAllTodos() ([]*domain.Todo, utils.Error)
	GetTodoByID(id int) (*domain.Todo, utils.Error)
}

type todoDomain struct{}

var TodoDomain todoDomainRepo = &todoDomain{}

func (t *todoDomain) CreateTodo(todo *domain.Todo) (*domain.Todo, utils.Error) {
	db := database.GetDB()

	err := db.Create(&todo).Error

	if err != nil {
		return nil, utils.InternalServerError("Something went wrong")
	}

	return todo, nil
}

func (t *todoDomain) UpdateTodo(input *domain.Todo, id int) (*domain.Todo, utils.Error) {
	db := database.GetDB()

	var todo domain.Todo
	
	err := db.First(&todo, id).Error

	if err != nil {
		return nil, utils.NotFound(fmt.Sprintf("Data with id %d not found", id))
	}

	todo.ID = id
	todo.Title = input.Title
	todo.Completed = input.Completed

	err = db.Save(&todo).Error

	if err != nil {
		return nil, utils.InternalServerError("Something went wrong")
	}
	
	return &todo, nil
}

func (t *todoDomain) DeleteTodo(id int) (utils.Error) {
	db := database.GetDB()

	var todo domain.Todo
	err := db.First(&todo, id).Error

	if err != nil {
		return utils.NotFound(fmt.Sprintf("Data with id %d not found", id))
	}

	err = db.Delete(&todo).Error

	if err != nil {
		return utils.InternalServerError("Something went wrong")
	}

	return nil
}

func (t *todoDomain) GetAllTodos() ([]*domain.Todo, utils.Error){
	db := database.GetDB()

	var todos []*domain.Todo

	err := db.Find(&todos).Error

	if err != nil {
		return nil, utils.InternalServerError("Something went wrong")
	}

	if len(todos) == 0 {
		return nil, utils.NotFound("The database is empty. Please add data to continue.")
	}

	return todos, nil
}

func (t *todoDomain) GetTodoByID(id int) (*domain.Todo, utils.Error){
	db := database.GetDB()
	var todo domain.Todo

	err := db.First(&todo, id).Error

	if err != nil {
		return nil, utils.NotFound(fmt.Sprintf("Data with id %d not found", id))
	}

	return &todo, nil
}