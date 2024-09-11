package models

type User struct {
	Id       int64  `json:"-"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
