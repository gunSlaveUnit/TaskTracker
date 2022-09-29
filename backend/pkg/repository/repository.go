package repository

import (
	"github.com/gunSlaveUnit/TaskTracker/pkg/db"
	"github.com/gunSlaveUnit/TaskTracker/pkg/entities"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
}

type Tasks interface {
}

type Repository struct {
	Authorization
	Tasks
}

func NewRepository(database *sqlx.DB) *Repository {
	return &Repository{Authorization: db.NewStorage(database)}
}
