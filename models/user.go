package models

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" db:"name" binding:"-"`
	Password string `json:"password" db:"hash_password" binding:"required"`
	Email    string `json:"email" db:"email" binding:"required"`
}
