package service

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/pkg"
)

func (service AuthService) Refresh(ctx context.Context, userId int) (entities.Token, error) {
	user, err := service.UserRepo.FindById(ctx, userId)

	if err != nil {
		return entities.Token{}, err
	}

	claim := pkg.PrepareAccessTokenClaims(user.Username, user.ID, user.Role)
	accessToken, err := pkg.GenerateToken(claim, pkg.TypeTokenAccess)

	if err != nil {
		return entities.Token{}, err
	}

	claim = pkg.PrepareRefreshTokenClaims(user.Username, user.ID)
	refreshToken, err := pkg.GenerateToken(claim, pkg.TypeTokenRefresh)

	if err != nil {
		return entities.Token{}, err
	}

	err = service.AuthRepo.Update(ctx, user.ID, refreshToken)

	if err != nil {
		return entities.Token{}, err
	}

	token := entities.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return token, nil
}
