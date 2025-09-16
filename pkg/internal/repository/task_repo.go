package repository

import "ToDoAmi/pkg/internal/models"

// interface for repository
type TaskRepository interface {
    GetTaskByID(id string) (*models.Task, error)
    CreateTask(task *models.Task) error
    UpdateTask(id string) error
    DeleteTask(id string) error
    SaveToFile() error
    LoadFromFile() ([]models.Task, error)
}