package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB(cfg *Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(cfg.DriverName, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	//db.Exec("CREATE TABLE Users\n(\n    id serial PRIMARY KEY,\n    name varchar(255) NOT NULL UNIQUE,\n    balance int\n);\nCREATE TABLE Quest\n(\n    id serial PRIMARY KEY,\n    name varchar(255) NOT NULL,\n    cost int CHECK (cost>0) NOT NULL,\n    times timestamp,\n    CONSTRAINT unique_quest UNIQUE (name, cost, times)\n);\nCREATE TABLE History\n(\n    id serial PRIMARY KEY,\n    questId int NOT NULL,\n    userId int NOT NULL,\n    status varchar(255) NOT NULL,\n    timeStart timestamp NOT NULL,\n    timeStop timestamp,\n    timeDeadLine timestamp,\n    FOREIGN KEY (questId) REFERENCES Quest(id),\n    FOREIGN KEY (userId) REFERENCES Users(id)\n);")
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
