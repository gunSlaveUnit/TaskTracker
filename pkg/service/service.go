package service

type Authorization interface {
}

type Games interface {
}

type Service struct {
	Authorization
	Games
}

func NewService() *Service {
	return &Service{}
}
