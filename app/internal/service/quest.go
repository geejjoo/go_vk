package service

import (
	"app/cmd/db"
	"app/internal/repository"
)

type QuestService struct {
	repo repository.Repository
}

func NewQuestService(repo repository.Repository) Quest {
	return &QuestService{repo: repo}
}

func (w *QuestService) QuestCreate(quest *db.Quest) error {
	err := w.repo.QuestCreate(quest)

	return err

}
