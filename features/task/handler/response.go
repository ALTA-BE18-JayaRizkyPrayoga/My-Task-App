package handler

type TaskResponse struct {
	ID        uint   `json:"id"`
	Status    string `json:"status"`
	ProjectID uint   `json:"project_id"`
}
