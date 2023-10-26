package service

import (
	"final-project-01/internal/domain"
	"final-project-01/internal/repository"
)

type todoServiceRepo interface {
	CreateTodo(*domain.Todo) (*domain.Todo, error)
	UpdateTodo(*domain.Todo, string) (*domain.Todo, error)
	DeleteTodo(string) (*domain.Todo, error)
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

func (t *todoService) UpdateTodo(todo *domain.Todo, id string) (*domain.Todo, error) {
	err := todo.Validate()

	if err != nil {
		return nil, err
	}

	todoResult, err := repository.TodoDomain.UpdateTodo(todo, id)

	if err != nil {
		return nil, err
	}

	return todoResult, nil
}

func (t *todoService) DeleteTodo(id string) (*domain.Todo, error) {
	todoResult, err := repository.TodoDomain.DeleteTodo(id)

	if err != nil {
		return nil, err
	}

	return todoResult, nil
}
