package auth

import "context"

func (repository AuthRepository) Save(ctx context.Context, userId int, refreshToken string) error {

	auth, err := repository.FindAuthByUserId(ctx, userId)

	if auth.UserId != 0 {
		repository.Update(ctx, userId, refreshToken)

		return nil
	}

	sql := `
		INSERT INTO
			auths(
				user_id,
				refresh_token)
			VALUES(?, ?)`

	_, err = repository.DB.ExecContext(ctx, sql,
		userId,
		refreshToken,
	)

	if err != nil {
		return err
	}

	return nil
}
