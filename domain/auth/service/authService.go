package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/auth"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/user"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/pkg"
)

type AuthService struct {
	UserRepo user.Repository
	AuthRepo auth.Repository
	DB       *sql.DB
}

func NewAuthService(authRepo auth.Repository, userRepo user.Repository, db *sql.DB) AuthService {
	return AuthService{
		AuthRepo: authRepo,
		UserRepo: userRepo,
		DB:       db,
	}
}

func (service AuthService) Login(ctx context.Context, login entities.Login) (entities.Token, error) {
	tx, err := service.DB.Begin()

	if err != nil {
		return entities.Token{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	user, err := service.UserRepo.FindByUsername(ctx, tx, login.Username)

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

	err = service.AuthRepo.Save(ctx, tx, user.ID, refreshToken)

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
	tx, err := service.DB.Begin()

	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	userId := ctx.Value("user_id").(int)
	_ = service.AuthRepo.Delete(ctx, tx, userId)

	return nil
}
func (service AuthService) Refresh(ctx context.Context) (entities.Token, error) {
	tx, err := service.DB.Begin()

	if err != nil {
		return entities.Token{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	userId := ctx.Value("user_id").(int)
	refreshToken := ctx.Value("refresh_token").(string)

	auth, err := service.AuthRepo.FindRefreshTokenByUserId(ctx, tx, userId)

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

	err = service.AuthRepo.Update(ctx, tx, userId, refreshToken)

	if err != nil {
		return entities.Token{}, err
	}

	token := entities.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return token, nil
}
