package models

type Task struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    IsCompleted bool   `json:"isCompleted"`
	Priority    string `json:"priority"`
}

type TaskRequest struct {
    ID          string `json:"id"`
    Title     string `json:"title"`
    Priority string `json:"priority"`
}

