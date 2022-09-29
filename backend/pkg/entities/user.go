package entities

type User struct {
	Id       int    `json:"-" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
