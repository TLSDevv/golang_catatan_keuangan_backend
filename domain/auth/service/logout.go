package service

import "context"

func (service AuthService) Logout(ctx context.Context) error {
	userId := ctx.Value("user_id").(int)
	_ = service.AuthRepo.Delete(ctx, userId)

	return nil
}
