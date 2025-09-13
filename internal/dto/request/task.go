package request

type TaskCreateRequest struct {
	Title       string `json:"title" binding:"required,min=1,max=255"`
	Description string `json:"description" binding:"max=1000"`
	ProjectID   uint   `json:"project_id" binding:"required"`
	Priority    uint   `json:"priority" binding:"min=1,max=3"`
	DueDate     string `json:"due_date,omitempty"`
}
