package report

import "github.com/and3407/go_reports/app/db/models"

type AddReportDto struct {
	TaskId int 
	WorkdayId int 
	TimeSpent int 
}

func AddReport(dto AddReportDto) (models.Report, bool) {
	report := models.Report{
		TimeSpent: dto.TimeSpent,
		TaskId: dto.TaskId,
		WorkdayId: dto.WorkdayId,
	}

	repository := GetReportRepository()

	return repository.AddReport(report)
}