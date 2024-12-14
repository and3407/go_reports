package repositories

import (
	"github.com/and3407/go_reports/app/db/models"
	"github.com/and3407/go_reports/app/db"
)

type ReportRepository struct {

}

func (repository ReportRepository) AddReport(report models.Report) (models.Report, bool)  {
	if result := db.Connect.Create(&report); result.Error != nil {
		return report, false
	}

	return report, true
}