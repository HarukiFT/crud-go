package service

import "github.com/jmoiron/sqlx"

type UserService struct {
	db *sqlx.DB
}

func (s *UserService) NewUser(name string, password string) bool {
	return true
}

func NewUserService(db *sqlx.DB) *UserService {
	userService := UserService{db: db}

	return &userService
}
