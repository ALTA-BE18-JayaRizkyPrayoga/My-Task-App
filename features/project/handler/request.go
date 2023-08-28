package handler

type ProjectRequest struct {
	Name        string `json:"name"`
	UserID      uint   `json:"user_id"`
	Description string `json:"description"`
}
