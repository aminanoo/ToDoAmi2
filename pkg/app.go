package internal

import (
 "fmt"
 "ToDoAmi/pkg/internal/service"
 "ToDoAmi/pkg/internal/repository"
 "ToDoAmi/pkg/internal/usecase"
 "ToDoAmi/pkg/internal/models"
 "context"
)

// app struct for main app
type App struct {
 TaskUseCase *usecase.TaskUseCase
 tasks []models.Task
}

// new app creation with inits
func NewApp() *App {
    TaskRepo := &repository.JSONTaskRepository{FilePath: "pkg/data/tasks.json"}
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

//func (a *App) Shutdown()
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
    // if err == nil {
    //     // Reload tasks from storage to sync local state
    //     tasks, loadErr := a.TaskUseCase.LoadData()
    //     if loadErr == nil {
    //         a.tasks = tasks
    //     }
    // }
    return err
}