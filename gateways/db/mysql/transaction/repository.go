package transaction

import (
	"database/sql"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/transaction"
)

var _ transaction.Repository = Repository{}

type Repository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
