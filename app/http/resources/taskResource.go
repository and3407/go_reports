package resources

import "github.com/and3407/go_reports/app/db/models"

type TaskResource struct {
	Task models.Task `json:"task"`
}