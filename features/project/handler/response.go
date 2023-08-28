package handler

type ProjectResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	UserID      uint   `json:"user_id"`
	Description string `json:"description"`
}
