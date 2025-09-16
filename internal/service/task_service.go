package service

import "ToDoAmi2/internal/models"

// logic interface
type TaskService interface {
    CreateTask(task models.TaskRequest) (*models.Task, error)
    UpdateTask(task *models.Task) error
    DeleteTask(id string) error
    CompleteTask(id string) error
    LoadData() ([]models.Task, error)
}