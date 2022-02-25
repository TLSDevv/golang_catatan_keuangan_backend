package auth

import (
	"database/sql"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return AuthRepository{
		DB: db,
	}
}
