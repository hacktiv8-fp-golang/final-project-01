package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// Todo represent the model for an todo
type Todo struct {
	ID int `gorm:"primaryKey" json:"id"`
	Title string `gorm:"not null;" json:"title" valid:"required~title is required"`
	Completed bool `gorm:"not null;default:false" json:"completed"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (todo *Todo) Validate() error {
	_, err := govalidator.ValidateStruct(todo)

	if err != nil {
		return err
	}

	return nil
}