package auth

import "context"

func (repository AuthRepository) Update(ctx context.Context, userId int, refreshToken string) error {
	sql := `
		UPDATE
			auths
		SET
			refresh_token=?
		WHERE
			user_id=?`

	_, err := repository.DB.ExecContext(ctx, sql,
		refreshToken,
		userId,
	)

	if err != nil {
		return err
	}

	return nil
}
