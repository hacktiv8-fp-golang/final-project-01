package repository

import (
	"final-project-01/internal/database"
	"final-project-01/internal/domain"
)

type todoDomainRepo interface {
	CreateTodo(*domain.Todo) (*domain.Todo, error)
	UpdateTodo(*domain.Todo, string) (*domain.Todo, error)
	DeleteTodo(string) (*domain.Todo, error)
	GetAllData() ([]*domain.Todo, error)
	GetDataByID(id int) (*domain.Todo, error)
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

func (t *todoDomain) UpdateTodo(input *domain.Todo, id string) (*domain.Todo, error) {
	db := database.GetDB()

	var todo domain.Todo
	err := db.First(&todo, id).Error

	if err != nil {
		return nil, err
	}

	db.Model(&todo).Updates(input)

	return &todo, nil
}

func (t *todoDomain) DeleteTodo(id string) (*domain.Todo, error) {
	db := database.GetDB()

	var todo domain.Todo
	err := db.First(&todo, id).Error

	if err != nil {
		return nil, err
	}

	db.Delete(&todo)

	return &todo, nil
}

func (t *todoDomain) GetAllData() ([]*domain.Todo, error){
	db := database.GetDB()

	var todos []*domain.Todo

	err := db.Find(&todos).Error

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (t *todoDomain) GetDataByID(id int) (*domain.Todo, error){
	db := database.GetDB()
	var todo domain.Todo

	err := db.First(&todo, id).Error 
	if err != nil {
		return nil, err
	}
	return &todo, nil
}