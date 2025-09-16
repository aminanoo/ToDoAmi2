package service

import "ToDoAmi/pkg/internal/models"

// logic interface
type TaskService interface {
    CreateTask(task models.TaskRequest) (*models.Task, error)
    UpdateTask(task *models.Task) error
    // GetFilteredTasks(filter models.FilterParams) ([]models.Task, error)
    // GetSortedTasks(sortBy string, ascending bool) ([]models.Task, error)
    DeleteTask(id string) error
    CompleteTask(id string) error
    LoadData() ([]models.Task, error)
}