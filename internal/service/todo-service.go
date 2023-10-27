package service

import (
	"final-project-01/internal/domain"
	"final-project-01/internal/repository"
)

type todoServiceRepo interface {
	CreateTodo(*domain.Todo) (*domain.Todo, error)
	UpdateTodo(*domain.Todo, int) (*domain.Todo, error)
	DeleteTodo(int) (error)
	GetAllTodos() (*[]domain.Todo, error)
	GetTodoByID(id int) (*domain.Todo, error)
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

func (t *todoService) UpdateTodo(todo *domain.Todo, id int) (*domain.Todo, error) {
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

func (t *todoService) DeleteTodo(id int) (error) {
	err := repository.TodoDomain.DeleteTodo(id)

	if err != nil {
		return err
	}

	return nil
}

func (t *todoService) GetAllTodos() (*[]domain.Todo, error){

	todoResult, err := repository.TodoDomain.GetAllTodos()

	if err != nil{
		return nil, err
	}

	todos := make([]domain.Todo, len(todoResult))
    for i, t := range todoResult {
        todos[i] = *t
    }

	return &todos, nil

}

func (t *todoService) GetTodoByID(id int ) (*domain.Todo, error){
	todoResult, err := repository.TodoDomain.GetTodoByID(id)

	if err != nil {
		return nil, err
	}

	return todoResult,nil
}