package models

type GetUserResponse struct {
	Id       uint64 `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type CreateUserResponse struct {
	GetUserResponse
}
