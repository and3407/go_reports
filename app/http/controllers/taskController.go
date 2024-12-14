package controllers

import (
	"net/http"

	"github.com/and3407/go_reports/app/domain/task"

	"github.com/gin-gonic/gin"
)

type TaskController struct {

}

func (controller TaskController) AddTask(ginContext *gin.Context) {
	dto := task.AddTaskDto{}

	if err := ginContext.BindJSON(&dto); err != nil {
		ginContext.AbortWithError(http.StatusBadRequest, err)
		return
	}

	task, ok := task.AddTask(dto)

	if (!ok) {
		ginContext.JSON(http.StatusInternalServerError, "")
	}

	ginContext.JSON(http.StatusCreated, task)
}