package model

type SignupForm struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginForm struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type User struct {
	Id       int    `json:"id" `
	Name     string `json:"name" `
	Email    string `json:"email" `
	Passhash string `json:"password"`
}
