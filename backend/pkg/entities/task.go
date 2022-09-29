package entities

type Task struct {
	Id          int    `json:"-"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Executor    int    `json:"executor" binding:"required"`
}
