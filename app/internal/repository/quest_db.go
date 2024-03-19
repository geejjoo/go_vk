package repository

import (
	"app/cmd/db"
	"github.com/jmoiron/sqlx"
	"strings"
)

type QuestDB struct {
	db            *sqlx.DB
	queryReplacer *strings.Replacer
}

func NewQuestDB(db *sqlx.DB, tableName string) Quest {
	return &QuestDB{
		db:            db,
		queryReplacer: strings.NewReplacer("{table}", tableName),
	}
}

func (w *QuestDB) QuestCreate(quest *db.Quest) error {
	query := w.queryReplacer.Replace("INSERT INTO {table} (name, cost, times) values ($1, $2, $3)")
	_, err := w.db.Exec(query, quest.Name, quest.Cost, quest.Time)
	if err != nil {
		return InvalidQuest
	}
	return nil
}
