package service

import (
	"github.com/HarukiFT/go-crud/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	db          *sqlx.DB
	UserService domain.UserService
}

func NewService(db *sqlx.DB) *Service {
	usecase := Service{
		db:          db,
		UserService: NewUserService(db),
	}

	return &usecase
}
