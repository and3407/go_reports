package requests

type CreateTaskRequest struct {
	Title string `json:"title" binding:"required"`
	Key string `json:"key" binding:"required"`
	Archived bool `json:"archived" binding:"boolean"`
}