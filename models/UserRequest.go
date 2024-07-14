package models

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Gmail    string `json:"gmail" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Gmail    string `json:"gmail"`
	Fullname string `json:"fullname"`
	Status   bool   `json:"status"`
}
