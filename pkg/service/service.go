package service

import "Ozinshe/pkg/repository"

type Authorization interface {
	CreateUser(user repository.SignupUser) (int, error)
	GenerateToken(email, password string) (string, error)
	// ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
