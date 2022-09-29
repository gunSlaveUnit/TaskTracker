package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Tasks interface {
}

type Repository struct {
	Authorization
	Tasks
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
