package transaction

import "context"

type Count struct {
	Total int `json:"total"`
}

func (r Repository) CheckUser(ctx context.Context, userID int) (bool, error) {
	var result Count

	query := `SELECT COUNT(*) as total FROM users WHERE id = ?`
	err := r.DB.QueryRowContext(ctx, query, userID).Scan(&result.Total)
	if err != nil {
		return false, err
	}

	return result.Total > 0, nil
}
