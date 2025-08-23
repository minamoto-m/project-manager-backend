package handlers

// Project represents a project
type Project struct {
	Id          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Color       *string `json:"color,omitempty"`
}

// ProjectListResponse represents a list of projects
type ProjectListResponse struct {
	Projects *[]Project `json:"projects,omitempty"`
}

// ProjectCreateRequest represents a project creation request
type ProjectCreateRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Color       *string `json:"color,omitempty"`
}

// ProjectUpdateRequest represents a project update request
type ProjectUpdateRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Color       *string `json:"color,omitempty"`
}

// ProjectResponse represents a single project response
type ProjectResponse struct {
	Project *Project `json:"project,omitempty"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Message *string                 `json:"message,omitempty"`
	Code    *string                 `json:"code,omitempty"`
	Details *map[string]interface{} `json:"details,omitempty"`
}
