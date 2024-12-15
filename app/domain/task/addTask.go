package task

import "github.com/and3407/go_reports/app/db/models"

type AddTaskDto struct {
	Title string 
	Key string
	Archived bool
}

func AddTask(dto AddTaskDto) (models.Task, bool) {
	task := models.Task{
		Title: dto.Title,
		Key: dto.Key,
		Archived: dto.Archived,
	}

	repository := GetTaskRepository()

	return repository.AddTask(task)
}