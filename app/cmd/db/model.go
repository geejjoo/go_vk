package db

import "time"

type User struct {
	ID      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Balance int    `json:"balance" db:"balance"`
}

type Quest struct {
	ID   int       `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`
	Cost int       `json:"cost" db:"cost"`
	Time time.Time `json:"time" db:"times"`
}

type History struct {
	ID           int       `json:"id" db:"id"`
	QuestID      int       `json:"questId" db:"questId"`
	UserID       int       `json:"userId" db:"userId"`
	TimeStart    time.Time `json:"time_start" db:"timeStart"`
	TimeStop     time.Time `json:"time_stop" db:"timeStop"`
	TimeDeadline time.Time `json:"time_deadline" db:"timeDeadline"`
	Status       string    `json:"status" db:"status"`
}

type UserInfo struct {
	ID      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Balance int    `json:"balance" db:"balance"`
	Quests  []int  `json:"quest" db:"questId"`
}
