package service

import (
	"context"
	"errors"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/auth"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/user"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/pkg"
)

type AuthService struct {
	UserRepo user.Repository
	AuthRepo auth.Repository
}

func NewAuthService(authRepo auth.Repository, userRepo user.Repository) AuthService {
	return AuthService{
		AuthRepo: authRepo,
		UserRepo: userRepo,
	}
}

func (service AuthService) Login(ctx context.Context, login entities.Login) (entities.Token, error) {
	user, err := service.UserRepo.FindByUsername(ctx, login.Username)

	if err != nil {
		return entities.Token{}, errors.New("Username Not Found")
	}

	if user.ID == 0 {
		return entities.Token{}, errors.New("Username Not Found")
	}

	if err := user.CheckPassword(login.Password); err != nil {
		return entities.Token{}, errors.New("Wrong Password")
	}

	// todo refactor pkg jwt
	claim := pkg.PrepareTokenClaims(login.Username, user.ID, pkg.TypeTokenAccess)
	accessToken, err := pkg.GenerateToken(claim, pkg.TypeTokenAccess)

	if err != nil {
		return entities.Token{}, err
	}

	claim = pkg.PrepareTokenClaims(login.Username, user.ID, pkg.TypeTokenRefresh)
	refreshToken, err := pkg.GenerateToken(claim, pkg.TypeTokenRefresh)

	if err != nil {
		return entities.Token{}, err
	}

	err = service.AuthRepo.Save(ctx, user.ID, refreshToken)

	if err != nil {
		return entities.Token{}, err
	}

	token := entities.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return token, nil
}
func (service AuthService) Logout(ctx context.Context) error {
	userId := ctx.Value("user_id").(int)
	_ = service.AuthRepo.Delete(ctx, userId)

	return nil
}
func (service AuthService) Refresh(ctx context.Context) (entities.Token, error) {
	userId := ctx.Value("user_id").(int)
	refreshToken := ctx.Value("refresh_token").(string)

	auth, err := service.AuthRepo.FindAuthByUserId(ctx, userId)

	if err != nil {
		return entities.Token{}, err
	}

	if refreshToken != auth.RefreshToken {
		return entities.Token{}, err
	}

	username := ctx.Value("username").(string)

	claim := pkg.PrepareTokenClaims(username, userId, pkg.TypeTokenAccess)
	accessToken, err := pkg.GenerateToken(claim, pkg.TypeTokenAccess)

	if err != nil {
		return entities.Token{}, err
	}

	claim = pkg.PrepareTokenClaims(username, userId, pkg.TypeTokenRefresh)
	refreshToken, err = pkg.GenerateToken(claim, pkg.TypeTokenRefresh)

	if err != nil {
		return entities.Token{}, err
	}

	err = service.AuthRepo.Update(ctx, userId, refreshToken)

	if err != nil {
		return entities.Token{}, err
	}

	token := entities.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return token, nil
}
