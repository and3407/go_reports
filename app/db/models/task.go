package models

import (
	"time"
)

type Task struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Key string `json:"key"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}