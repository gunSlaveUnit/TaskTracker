package service

import "github.com/gunSlaveUnit/TaskTracker/pkg/repository"

type Authorization interface {
}

type Tasks interface {
}

type Service struct {
	Authorization
	Tasks
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
