package service

import "errors"

var DatabaseError = errors.New("Database error")

var UserCreateError = errors.New("Cannot create user")

var UserNotFoundError = errors.New("Cannot get info about user")

var InvalidInfo = errors.New("Deadline is expired")

var InvalidQuest = errors.New("Cannot create quest")
