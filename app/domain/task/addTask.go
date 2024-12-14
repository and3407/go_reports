package task

import "github.com/and3407/go_reports/app/db/models"

type AddTaskDto struct {
	Title string `json:"title"`
	Key string `json:"key"`
}

func AddTask(dto AddTaskDto) (models.Task, bool) {
	task := models.Task{
		Title: dto.Title,
		Key: dto.Key,
	}

	repository := GetTaskRepository()

	return repository.AddTask(task)
}