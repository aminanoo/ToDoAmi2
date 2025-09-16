package internal

import (
 "fmt"
 "ToDoAmi2/internal/service"
 "ToDoAmi2/internal/repository"
 "ToDoAmi2/internal/usecase"
 "ToDoAmi2/internal/models"
 "context"
)

// app struct for main app
type App struct {
 TaskUseCase *usecase.TaskUseCase
 tasks []models.Task
}

// new app creation with inits
func NewApp() *App {
    TaskRepo := &repository.JSONTaskRepository{FilePath: "internal/data/tasks.json"}
    TaskService := &service.TaskServiceImpl{TaskRepo: *TaskRepo}
    TaskUseCase := &usecase.TaskUseCase{TaskService: *TaskService}
    return &App{TaskUseCase: TaskUseCase,}
}

// starting app and load from file
func (a *App) Startup(ctx context.Context) {
    tasks, err := a.TaskUseCase.LoadData()
    a.tasks = tasks
    if err != nil {
        fmt.Println("Error loading data:", err)
    }
}

func (a *App) GetTasks() []models.Task {
	return a.tasks
}

// HandleAddTask for adding tasks
func (a *App) CreateTask(task models.TaskRequest) (*models.Task, error) {
    newTask, err := a.TaskUseCase.HandleAddTask(task)
    if err != nil {
        return nil, err
    }
    return newTask, nil
}

func (a *App) DeleteTask(id string) error {
    err := a.TaskUseCase.TaskService.DeleteTask(id)
    return err
}

func (a *App) UpdateTask(id string) (error) {
    err := a.TaskUseCase.TaskService.CompleteTask(id)
    return err
}