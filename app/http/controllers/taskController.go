package controllers

import (
	"net/http"

	"github.com/and3407/go_reports/app/domain/task"
	"github.com/and3407/go_reports/app/http/requests/task"

	"github.com/gin-gonic/gin"
)

type TaskController struct {

}

func (controller TaskController) AddTask(ginContext *gin.Context) {
	var request requests.CreateTaskRequest
	
	if err := ginContext.ShouldBindJSON(&request); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	dto := task.AddTaskDto{
		Title: request.Title,
		Key: request.Key,
		Archived: request.Archived,
	}

	task, ok := task.AddTask(dto)

	if (!ok) {
		ginContext.JSON(http.StatusInternalServerError, "")
	}

	ginContext.JSON(http.StatusCreated, task)
}