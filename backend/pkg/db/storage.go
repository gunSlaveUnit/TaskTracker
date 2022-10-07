package db

import (
	"fmt"
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
	rawQuery := fmt.Sprintf("INSERT INTO users (name, password) values (%s, %s) RETURNING id;", user.Name, user.Password)

	var id int
	insertedData := s.db.QueryRow(rawQuery)
	if err := insertedData.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}
