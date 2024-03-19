package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	User
	Quest
}

func NewRepository(
	db *sqlx.DB,
	userTableName string,
	questTableName string,
) *Repository {
	return &Repository{
		NewUserDB(db, userTableName),
		NewQuestDB(db, questTableName),
	}
}
