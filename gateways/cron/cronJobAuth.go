package cron

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/auth"
	"github.com/sirupsen/logrus"
)

type CronJobAuth struct {
	AuthService auth.Service
}

func NewCronJobAuth(authService auth.Service) *CronJobAuth {
	return &CronJobAuth{
		AuthService: authService,
	}
}

func (cron CronJobAuth) DeletedAuthNotValid(ctx context.Context) {
	logrus.Info("CronJob Deleted Auth Not Valid Run")

	totalDeletedAuth, err := cron.AuthService.DeletedAuthNotValid(ctx)
	if err != nil {
		logrus.Error(err.Error())
	}

	logrus.WithFields(logrus.Fields{
		"msg":           "Done",
		"total deleted": totalDeletedAuth,
	}).Info("Cronjob Success")
}
