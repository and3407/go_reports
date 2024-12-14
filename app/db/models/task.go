package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model 
	Title string
	Key string `gorm:"index:idx_key_unique,unique;not null"`
}