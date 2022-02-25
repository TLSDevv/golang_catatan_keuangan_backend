package user

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (r UserRepository) CheckUser(ctx context.Context, userID int) (bool, error) {
	var result entities.CountResult

	query := `SELECT COUNT(*) as total FROM users WHERE id = ?`
	err := r.DB.QueryRowContext(ctx, query, userID).Scan(&result.Total)
	if err != nil {
		return false, err
	}

	return result.Total > 0, nil
}
