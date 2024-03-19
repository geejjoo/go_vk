package service

import "app/internal/repository"

type Service struct {
	User
	Quest
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		NewUserService(*repos),
		NewQuestService(*repos),
	}
}
