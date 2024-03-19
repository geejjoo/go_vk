package repository

import (
	"app/cmd/db"
)

type User interface {
	UserCreate(name string) (int, error)
	UserInfo(id int) (db.UserInfo, error)
	UserQuest(history *db.History) error
}

type Quest interface {
	QuestCreate(quest *db.Quest) error
}
