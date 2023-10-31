package service

import (
	"github.com/hacktiv8-fp-golang/final-project-01/internal/domain"
	"github.com/hacktiv8-fp-golang/final-project-01/internal/repository"
	"github.com/hacktiv8-fp-golang/final-project-01/internal/utils"
)

type todoServiceRepo interface {
	CreateTodo(*domain.Todo) (*domain.Todo, utils.Error)
	UpdateTodo(*domain.Todo, int) (*domain.Todo, utils.Error)
	DeleteTodo(int) (utils.Error)
	GetAllTodos() (*[]domain.Todo, utils.Error)
	GetTodoByID(id int) (*domain.Todo, utils.Error)
}

type todoService struct{}

var TodoService todoServiceRepo = &todoService{}

func (t *todoService) CreateTodo(todo *domain.Todo) (*domain.Todo, utils.Error) {
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

func (t *todoService) UpdateTodo(todo *domain.Todo, id int) (*domain.Todo, utils.Error) {
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

func (t *todoService) DeleteTodo(id int) (utils.Error) {
	err := repository.TodoDomain.DeleteTodo(id)

	if err != nil {
		return err
	}

	return nil
}

func (t *todoService) GetAllTodos() (*[]domain.Todo, utils.Error){

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

func (t *todoService) GetTodoByID(id int ) (*domain.Todo, utils.Error){
	todoResult, err := repository.TodoDomain.GetTodoByID(id)

	if err != nil {
		return nil, err
	}

	return todoResult,nil
}