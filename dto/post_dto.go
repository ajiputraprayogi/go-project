package dto

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type PostResponse struct {
	ID      uint         `json:"id"`
	Title   string       `json:"title"`
	Content string       `json:"content"`
	UserID  uint         `json:"user_id"`
	User    UserResponse `json:"user"`
}
