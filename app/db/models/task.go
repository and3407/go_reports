package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Key string `json:"key"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}