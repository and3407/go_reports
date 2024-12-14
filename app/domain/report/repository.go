package report

import (
	"github.com/and3407/go_reports/app/db/models"
	"github.com/and3407/go_reports/app/db/repositories"
)

type IRepository interface {
	AddReport(report models.Report) (models.Report, bool)
}

func GetReportRepository() IRepository {
	return repositories.ReportRepository{}
}