package helper

import (
	"database/sql"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func SqlRowsToTransactions(rows *sql.Rows) []entities.Transaction {
	result := []entities.Transaction{}
	if rows.Next() != true {
		return result
	}

	for rows.Next() {
		transaction := entities.Transaction{}

		err := rows.Scan(
			&transaction.ID,
			&transaction.TransactionName,
			&transaction.Category,
			&transaction.TransactionType,
			&transaction.Amount,
			&transaction.TransactionAt,
			&transaction.CreatedAt,
		)

		if err == nil {
			return result
		}

		result = append(result, transaction)
	}

	return result
}
