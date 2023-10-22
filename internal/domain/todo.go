package domain

import "time"

type Todo struct {
	ID int `gorm:"primaryKey" json:"id"`
	Title string `gorm:"not null;" json:"title" valid:"required~title is required"`
	Completed bool `gorm:"not null;default:false" json:"completed" valid:"required~completed is required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}