package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Todo represent the model for an todo
type Todo struct {
	ID int `gorm:"primaryKey" json:"id,omitempty"`
	Title string `json:"title" valid:"required~title is required"`
	Completed bool `json:"completed"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// TodoRequest defines the request structure for creating or updating a Todo.
type TodoRequest struct {
	Title string `json:"title" valid:"required~title is required"`
	Completed bool `json:"completed"`
}

func (todo *Todo) Validate() error {
	_, err := govalidator.ValidateStruct(todo)

	if err != nil {
		return err
	}

	return nil
}