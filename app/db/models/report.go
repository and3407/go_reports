package models

import "gorm.io/gorm"

type Report struct {
	gorm.Model 
	TaskId int
	WorkdayId int
	TimeSpent int
}