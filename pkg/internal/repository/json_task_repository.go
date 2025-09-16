package repository

import (
    "encoding/json"
    "fmt"
    "os"
    "ToDoAmi/pkg/internal/models"
    "io"
    "errors"
)

type JSONTaskRepository struct {
    tasks map[string]models.Task
    FilePath string
}

func (repo *JSONTaskRepository) LoadFromFile() ([]models.Task, error) {
    file, err := os.Open(repo.FilePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    repo.tasks = make(map[string]models.Task)
    err = decoder.Decode(&repo.tasks)
    if err != nil {
        if errors.Is(err, io.EOF) {
			return []models.Task{}, nil
		}
        return nil, err
    }
    var taskList []models.Task

    for _, task := range repo.tasks {
        taskList = append(taskList, task)
    }

    return taskList, nil
}

func (repo *JSONTaskRepository) GetTaskByID(id string) (*models.Task, error) {
    task, ok := repo.tasks[id]
    if ok {
		return &task, nil
	}
	return nil, fmt.Errorf("task not found")
}

func (repo *JSONTaskRepository) CreateTask(task *models.Task) error {
	repo.tasks[task.ID] = *task
	return repo.SaveToFile()
}

func (repo *JSONTaskRepository) UpdateTask(id string) error {
    task, exists := repo.tasks[id]
    if !exists {
     return fmt.Errorf("task not found")
    }
    task.IsCompleted = !task.IsCompleted
    repo.tasks[id] = task
    return repo.SaveToFile()
}

func (repo *JSONTaskRepository) DeleteTask(id string) error {
    _, exists := repo.tasks[id]
    if !exists {
        return fmt.Errorf("task not found")
    }
    delete(repo.tasks, id)
    return repo.SaveToFile()
}

// not optimal but works
func (repo *JSONTaskRepository) SaveToFile() error {
    data, err := json.MarshalIndent(repo.tasks, "", "  ")
    if err != nil {
        return err
    }

    err = os.WriteFile(repo.FilePath, data, 0644)
    if err != nil {
        return err
    }

    return nil
}
