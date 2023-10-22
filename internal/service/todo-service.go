package service

import (
	"final-project-01/internal/domain"
	"final-project-01/internal/repository"
)

type todoServiceRepo interface {
	CreateTodo(*domain.Todo) (*domain.Todo, error)
}

type todoService struct{}

var TodoService todoServiceRepo = &todoService{}

func (t *todoService) CreateTodo(todo *domain.Todo) (*domain.Todo, error) {
	err := todo.Validate()

	if err != nil {
		return nil, err
	}

	todoResult, err := repository.TodoDomain.CreateTodo(todo)

	if err != nil {
		return nil, err
	}

	return todoResult, nil
}