package db

import (
	"github.com/gunSlaveUnit/TaskTracker/pkg/entities"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreateUser(user entities.User) (int, error) {
	return -1, nil
}
