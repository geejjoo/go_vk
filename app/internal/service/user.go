package service

import (
	"app/cmd/db"
	"app/internal/repository"
)

type UserService struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) User {
	return &UserService{repo: repo}
}

func (w *UserService) UserCreate(name string) (int, error) {
	id, err := w.repo.UserCreate(name)
	if err != nil {
		return 0, UserCreateError
	}
	return id, nil

}

func (w *UserService) UserInfo(id int) (db.UserInfo, error) {
	userList, err := w.repo.UserInfo(id)
	if err != nil {
		return db.UserInfo{}, UserNotFoundError
	}

	return userList, err

}

func (w *UserService) UserQuest(history *db.History) error {
	err := w.repo.UserQuest(history)

	return err

}
