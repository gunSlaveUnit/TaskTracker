package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Games interface {
}

type Repository struct {
	Authorization
	Games
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
