package task

import (
	"github.com/and3407/go_reports/app/db/models"
	"github.com/and3407/go_reports/app/db/repositories"
)

type ITaskRepository interface {
	AddTask(task models.Task) (models.Task, bool)
}

func GetTaskRepository() ITaskRepository {
	return repositories.TaskRepository{}
}