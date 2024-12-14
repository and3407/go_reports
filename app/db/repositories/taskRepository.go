package repositories

import (
	"github.com/and3407/go_reports/app/db/models"
	"github.com/and3407/go_reports/app/db"
)

type TaskRepository struct {

}

func (repository TaskRepository) AddTask(task models.Task) (models.Task, bool)  {
	if result := db.Connect.Create(&task); result.Error != nil {
		return task, false
	}

	return task, true
}