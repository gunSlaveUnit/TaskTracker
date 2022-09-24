package entities

type Task struct {
	Id          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Executor    int    `json:"executor"`
}
