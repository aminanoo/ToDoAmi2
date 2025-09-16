package service

import (
    "ToDoAmi/pkg/internal/models"
    "ToDoAmi/pkg/internal/repository"
    //"time"
    "fmt"
)

// implementing task service
type TaskServiceImpl struct {
    TaskRepo repository.JSONTaskRepository
}

func (service *TaskServiceImpl) CreateTask(task models.TaskRequest) (*models.Task, error) {
    newTask := models.Task{
        ID:        task.ID,
        Title:     task.Title,
        //CreatedAt: time.Now().Format(time.RFC3339),
        //DueDate:   task.DueDate,
		Priority:  task.Priority,
		IsCompleted: false,
    }

    err := service.TaskRepo.CreateTask(&newTask)
    if err != nil {
        return nil, err
    }
    return &newTask, nil
}

// changes the IsComplete to !(IsComplete)
func (service *TaskServiceImpl) CompleteTask(id string) error {
    err := service.TaskRepo.UpdateTask(id)
    return err
}

// func generateUniqueID() string {
//     return fmt.Sprintf("%d", time.Now().UnixNano())
// }

func (service *TaskServiceImpl) LoadData() ([]models.Task, error) {
	data, err := service.TaskRepo.LoadFromFile()
	if err != nil {
		return nil, fmt.Errorf("error loading data: %v", err)
	}
	return data, nil
}

func (service *TaskServiceImpl) DeleteTask(id string) error {
    err := service.TaskRepo.DeleteTask(id)
    return err
}