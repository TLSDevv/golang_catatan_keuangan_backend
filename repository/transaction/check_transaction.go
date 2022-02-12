package transaction

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (r Repository) CheckTransactionByID(ctx context.Context, tID int) (bool, error) {
	var result entities.CountResult

	query := `SELECT COUNT(*) as total FROM transactions WHERE id = ?`
	err := r.DB.QueryRowContext(ctx, query, tID).Scan(&result.Total)
	if err != nil {
		return false, err
	}

	return result.Total > 0, nil
}
