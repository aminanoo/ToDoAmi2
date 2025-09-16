package usecase

import (
    "ToDoAmi/pkg/internal/models"
    "ToDoAmi/pkg/internal/service"
    "strings"
    "fmt"
)

type TaskUseCase struct {
    TaskService service.TaskServiceImpl
}

func (useCase *TaskUseCase) HandleAddTask(task models.TaskRequest) (*models.Task, error) {
    task.Title = strings.TrimSpace(task.Title)
    if task.Title == "" {
        return nil, fmt.Errorf("can't add task with empty title")
    }
    if task.Priority == "" {
        task.Priority = "low"
    }
    newTask, err := useCase.TaskService.CreateTask(task)
    if err != nil {
        return nil, err
    }

    return &models.Task{
        ID:        newTask.ID,
        Title:     newTask.Title,
        IsCompleted: newTask.IsCompleted,
        //CreatedAt: newTask.CreatedAt,
        //DueDate:   newTask.DueDate,
        Priority:  newTask.Priority,
    }, nil
}

func (useCase *TaskUseCase) LoadData() ([]models.Task, error) {
	data, err := useCase.TaskService.LoadData()
	if err != nil {
		return nil, err
	}
	return data, nil
}
