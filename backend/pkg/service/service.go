package service

import "github.com/gunSlaveUnit/TaskTracker/pkg/repository"

type Authorization interface {
}

type Games interface {
}

type Service struct {
	Authorization
	Games
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
