CREATE TABLE Users
(
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL UNIQUE,
    balance int
);
CREATE TABLE Quest
(
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL,
    cost int CHECK (cost>0) NOT NULL,
    times timestamp NOT NULL,
    CONSTRAINT unique_quest UNIQUE (name, cost, times)
);
CREATE TABLE History
(
    id serial PRIMARY KEY,
    questId int NOT NULL,
    userId int NOT NULL,
    status varchar(255) NOT NULL,
    timeStart timestamp NOT NULL,
    timeStop timestamp,
    timeDeadLine timestamp,
    FOREIGN KEY (questId) REFERENCES Quest(id),
    FOREIGN KEY (userId) REFERENCES Users(id)
);