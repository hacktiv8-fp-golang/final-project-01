package repository

import (
	"github.com/hacktiv8-fp-golang/final-project-01/internal/database"
	"github.com/hacktiv8-fp-golang/final-project-01/internal/domain"
	"github.com/hacktiv8-fp-golang/final-project-01/internal/utils"
)

type todoDomainRepo interface {
	CreateTodo(*domain.Todo) (*domain.Todo, utils.Error)
	UpdateTodo(*domain.TodoUpdate, int) (*domain.Todo, utils.Error)
	DeleteTodo(int) utils.Error
	GetAllTodos() ([]*domain.Todo, utils.Error)
	GetTodoByID(int) (*domain.Todo, utils.Error)
}

type todoDomain struct{}

var TodoDomain todoDomainRepo = &todoDomain{}

func (t *todoDomain) CreateTodo(newTodo *domain.Todo) (*domain.Todo, utils.Error) {
	db := database.GetDB()

	err := db.Create(&newTodo).Error

	if err != nil {
		return nil, utils.ParseError(err)
	}

	return newTodo, nil
}

func (t *todoDomain) UpdateTodo(updatedTodo *domain.TodoUpdate, id int) (*domain.Todo, utils.Error) {
	db := database.GetDB()

	var todo domain.Todo

	err := db.First(&todo, id).Error

	if err != nil {
		return nil, utils.ParseError(err)
	}

	todo.ID = id
	todo.Title = updatedTodo.Title
	todo.Completed = updatedTodo.Completed

	err = db.Save(&todo).Error

	if err != nil {
		return nil, utils.ParseError(err)
	}

	return &todo, nil
}

func (t *todoDomain) DeleteTodo(id int) utils.Error {
	db := database.GetDB()

	var todo domain.Todo
	err := db.First(&todo, id).Error

	if err != nil {
		return utils.ParseError(err)
	}

	err = db.Delete(&todo).Error

	if err != nil {
		return utils.ParseError(err)
	}

	return nil
}

func (t *todoDomain) GetAllTodos() ([]*domain.Todo, utils.Error){
	db := database.GetDB()

	var todos []*domain.Todo

	err := db.Find(&todos).Error

	if err != nil {
		return nil, utils.ParseError(err)
	}

	return todos, nil
}

func (t *todoDomain) GetTodoByID(id int) (*domain.Todo, utils.Error){
	db := database.GetDB()
	var todo domain.Todo

	err := db.First(&todo, id).Error

	if err != nil {
		return nil, utils.ParseError(err)
	}

	return &todo, nil
}
