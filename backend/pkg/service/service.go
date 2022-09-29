package service

import (
	"github.com/gunSlaveUnit/TaskTracker/pkg/entities"
	"github.com/gunSlaveUnit/TaskTracker/pkg/repository"
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
}

type Tasks interface {
}

type Service struct {
	Authorization
	Tasks
}

func NewService(repository *repository.Repository) *Service {
	return &Service{Authorization: NewAuthorizationService(repository.Authorization)}
}
