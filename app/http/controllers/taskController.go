package controllers

import (
	"net/http"

	"github.com/and3407/go_reports/app/domain/task"

	"github.com/gin-gonic/gin"
)

type TaskController struct {

}

type CreateTaskRequest struct {
	Title string `json:"title" binding:"required"`
	Key string `json:"key" binding:"required"`
	Archived bool `json:"archived" binding:"boolean"`
}

func (controller TaskController) AddTask(ginContext *gin.Context) {
	var request CreateTaskRequest
	
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