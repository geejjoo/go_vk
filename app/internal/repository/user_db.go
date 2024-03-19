package repository

import (
	"app/cmd/db"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type UserDB struct {
	db            *sqlx.DB
	queryReplacer *strings.Replacer
}

func NewUserDB(db *sqlx.DB, tableName string) User {
	return &UserDB{
		db:            db,
		queryReplacer: strings.NewReplacer("{table}", tableName),
	}
}

func (w *UserDB) UserCreate(name string) (int, error) {
	var id int
	query := w.queryReplacer.Replace("INSERT INTO {table} (balance, name) values (0,$1) RETURNING id")
	err := w.db.QueryRow(query, name).Scan(&id)
	if err != nil {
		return 0, InvalidCreate
	}
	return id, nil
}

func (w *UserDB) UserInfo(id int) (db.UserInfo, error) {
	var userInfo db.UserInfo

	tx, err := w.db.Beginx()
	if err != nil {
		return db.UserInfo{}, err
	}
	defer tx.Rollback()

	var balance int
	balanceQuery := w.queryReplacer.Replace("SELECT balance FROM Users WHERE id = $1")
	err = tx.Get(&balance, balanceQuery, id)
	if err != nil {
		return db.UserInfo{}, err
	}

	userInfo.ID = id
	userInfo.Balance = balance

	var completedQuestIDs []int
	questsQuery := w.queryReplacer.Replace("SELECT questId FROM History WHERE userId = $1")
	err = tx.Select(&completedQuestIDs, questsQuery, id)
	if err != nil && err != sql.ErrNoRows {
		return db.UserInfo{}, err
	}

	userInfo.Quests = completedQuestIDs

	if err := tx.Commit(); err != nil {
		return db.UserInfo{}, err
	}

	return userInfo, nil
}

func (w *UserDB) UserQuest(history *db.History) error {
	tx, err := w.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var timeStart sql.NullTime

	err = tx.QueryRow("SELECT timeStart FROM History WHERE questId = $1 AND userId = $2", history.QuestID, history.UserID).Scan(&timeStart)
	if err != nil {
		if err == sql.ErrNoRows {
			_, err = tx.Exec("INSERT INTO History (questId, userId, timeStart, status) VALUES ($1, $2, $3, 'run')", history.QuestID, history.UserID, time.Now())
			if err != nil {
				fmt.Println(err, "1")
				return err
			}
			fmt.Println(err, 2)
			return nil
		}
		fmt.Println(err, 3)
		return DatabaseError
	}

	// Проверяем, выполнено ли задание пользователем
	var isCompleted bool
	err = tx.QueryRow("SELECT EXISTS (SELECT 1 FROM History WHERE questId = $1 AND userId = $2 and status='completed')", history.QuestID, history.UserID).Scan(&isCompleted)
	fmt.Println(err)
	if err != nil {
		return err
	}
	if isCompleted {
		return errors.New("Task already completed")
	}

	// Получаем стоимость задания и ID пользователя
	quest := db.Quest{}
	err = tx.QueryRow("SELECT cost FROM Quest WHERE id = $1", history.QuestID).Scan(&quest.Cost)
	fmt.Println(err)
	if err != nil {
		return err
	}

	// Обновляем баланс пользователя
	_, err = tx.Exec("UPDATE Users SET balance = balance + $1 WHERE id = $2", quest.Cost, history.UserID)
	fmt.Println(quest.Cost)
	if err != nil {
		return err
	}

	// Записываем выполненное задание в историю
	_, err = tx.Exec("UPDATE History Set status = 'completed', timeStop = $1 where questId = $2 and userId = $3 and status= 'run'", time.Now(), history.QuestID, history.UserID)
	if err != nil {
		return err
	}

	return nil
}
