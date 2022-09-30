package models

type Task struct {
	UserId   string `json:"user_id"`
	Title    string `json:"title"`
	Comment  string `json:"comment"`
	Color    string `json:"color"`
	WithDate string `json:"with_date"`
}

type TaskResp struct {
	TaskId    string `json:"task_id"`
	UserId    string `json:"user_id"`
	Title     string `json:"title"`
	Comment   string `json:"comment"`
	Color     string `json:"color"`
	WithDate  string `json:"with_date"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TaskUpdate struct {
	UserId   string `json:"user_id"`
	Title    string `json:"title"`
	Comment  string `json:"comment"`
	Color    string `json:"color"`
	WithDate string `json:"with_date"`
}

type TasksList struct {
	Tasks []TaskResp `json:"tasks"`
	Count int64      `json:"count"`
}
