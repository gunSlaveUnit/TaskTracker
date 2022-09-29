package service

import (
	"github.com/gunSlaveUnit/TaskTracker/pkg/entities"
	"github.com/gunSlaveUnit/TaskTracker/pkg/repository"
)

type AuthorizationService struct {
	r repository.Authorization
}

func NewAuthorizationService(repo repository.Authorization) *AuthorizationService {
	return &AuthorizationService{r: repo}
}

func (authService AuthorizationService) CreateUser(user entities.User) (int, error) {
	return authService.r.CreateUser(user)
}
