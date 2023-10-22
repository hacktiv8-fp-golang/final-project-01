package repository

import (
	"final-project-01/internal/database"
	"final-project-01/internal/domain"
)

type todoDomainRepo interface {
	CreateTodo(*domain.Todo) (*domain.Todo, error)
}

type todoDomain struct{}

var TodoDomain todoDomainRepo = &todoDomain{}

func (t *todoDomain) CreateTodo(todo *domain.Todo) (*domain.Todo, error) {
	db := database.GetDB()

	err := db.Create(&todo).Error

	if err != nil {
		return nil, err
	}

	return todo, nil
}