package repository

import "errors"

var DatabaseError = errors.New("Database error")

var InvalidCreate = errors.New("Cannot create user")

var InvalidInfo = errors.New("Cannot get info about user")

var QuestDedlineExpired = errors.New("Deadline is expired")

var InvalidQuest = errors.New("Cannot create quest")
