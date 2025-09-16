package models

//import "time"

type Task struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    IsCompleted bool   `json:"isCompleted"`
    //CreatedAt   string `json:"createdAt"`
    //DueDate     *time.Time `json:"dueDate"`
	Priority    string `json:"priority"`
}

type TaskRequest struct {
    ID          string `json:"id"`
    Title     string `json:"title"`
    //DueDate   *time.Time `json:"dueDate"`
    Priority string `json:"priority"` // low, medium, high
}

