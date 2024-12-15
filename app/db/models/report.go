package models

import (
	"time"
)

type Report struct {
	Id uint `json:"id"`
	TaskId int `json:"taskId"`
	WorkdayId int `json:"workdayId"`
	TimeSpent int `json:"timeSpent"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}