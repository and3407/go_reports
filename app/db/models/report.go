package models

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	Id int `json:"id"`
	TaskId int `json:"taskId"`
	WorkdayId int `json:"workdayId"`
	TimeSpent int `json:"timeSpent"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}