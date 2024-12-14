package controllers

import (
	"net/http"

	"github.com/and3407/go_reports/app/domain/report"
	"github.com/gin-gonic/gin"
)

type ReportController struct {

}

func (controller ReportController) AddReport(ginContext *gin.Context) {
	dto := report.AddReportDto{}

	if err := ginContext.BindJSON(&dto); err != nil {
		ginContext.AbortWithError(http.StatusBadRequest, err)
		return
	}

	report, ok := report.AddReport(dto)

	if (!ok) {
		ginContext.JSON(http.StatusInternalServerError, "")
	}

	ginContext.JSON(http.StatusCreated, report)
}