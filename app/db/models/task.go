package models

import (
	"time"
)

type Task struct {
	Id uint `json:"id"`
	Title string `json:"title"`
	Key string `json:"key"`
	Archived bool `json:"archived"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}