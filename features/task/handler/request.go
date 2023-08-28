package handler

type TaskRequest struct {
	Status    string `json:"status"`
	ProjectID uint   `json:"project_id"`
}
