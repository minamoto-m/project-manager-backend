package response

import "time"

type TaskResponse struct {
	ID          uint      `json:"id"`
	ProjectID   uint      `json:"project_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    uint      `json:"priority"`
	DueDate     string    `json:"due_date"`
	CreatorID   uint      `json:"creator_id"`
	AssigneeID  uint      `json:"assignee_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskSummaryResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Priority uint   `json:"priority"`
	DueDate  string `json:"due_date"`
}

type ListMeta struct {
	Page    int  `json:"page"`
	PerPage int  `json:"per_page"`
	Total   int  `json:"total"`
	HasNext bool `json:"has_next"`
}

type ListResponse[T any] struct {
	Items []T      `json:"items"`
	Meta  ListMeta `json:"meta"`
}
