package repository

type Authorization interface {
}

type Games interface {
}

type Repository struct {
	Authorization
	Games
}

func NewRepository() *Repository {
	return &Repository{}
}
