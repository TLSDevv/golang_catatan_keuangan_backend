package service

import (
	"context"
	"fmt"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/pkg"
	"github.com/sirupsen/logrus"
)

func (service AuthService) DeletedAuthNotValid(ctx context.Context) (int, error) {
	users, err := service.AuthRepo.GetAllAuth(ctx)

	if err != nil {
		logrus.Error(err.Error())
		return 0, err
	}

	var count int

	if len(users) == 0 {
		return 0, err
	}

	for i, _ := range users {
		valid := pkg.RefreshTokenIsValid(users[i].RefreshToken)

		if !valid {
			err = service.AuthRepo.Delete(ctx, users[i].UserId)

			if err != nil {
				logrus.Error(err.Error())
				return 0, err
			}
			count += 1
		}
	}

	return count, nil
}
