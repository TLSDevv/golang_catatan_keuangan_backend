package service

import (
	"context"
	"errors"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/pkg"
)

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
	claim := pkg.PrepareAccessTokenClaims(login.Username, user.ID, user.Role)
	accessToken, err := pkg.GenerateToken(claim, pkg.TypeTokenAccess)

	if err != nil {
		return entities.Token{}, err
	}

	claim = pkg.PrepareRefreshTokenClaims(login.Username, user.ID)
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
