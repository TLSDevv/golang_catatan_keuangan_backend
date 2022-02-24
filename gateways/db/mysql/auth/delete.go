package auth

import "context"

func (repository AuthRepository) Delete(ctx context.Context, userId int) error {
	sql := `
		DELETE FROM
			auths
		WHERE
			user_id=?`

	_, err := repository.DB.ExecContext(ctx, sql,
		userId,
	)

	if err != nil {
		return err
	}

	return nil
}
