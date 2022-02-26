package user

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/helper"
	"github.com/sirupsen/logrus"
)

func (h UserHandler) Create(w http.ResponseWriter, r *http.Request) {

	var reqBody RequestUser

	err := util.Decode(r, &reqBody)

	if err != nil {
		util.SendNoData(w, http.StatusBadRequest, err.Error())
		logrus.WithFields(logrus.Fields{
			"domain":  "User",
			"handler": "Create User",
			"err":     err.Error(),
		}).Error("Decode")
		return
	}

	var validationErrPayload ValidationErrorResponse
	err = h.validator.Validate(reqBody, &validationErrPayload)

	if err != nil {
		util.SendWithData(w, http.StatusUnprocessableEntity, err.Error(), validationErrPayload)
		logrus.WithFields(logrus.Fields{
			"domain":  "User",
			"handler": "Create User",
			"err":     err.Error(),
		}).Error("Validate")
		return
	}

	err = h.usService.Create(r.Context(), entities.User{
		Username: reqBody.Username,
		Email:    reqBody.Email,
		Password: reqBody.Password,
		Fullname: reqBody.Fullname,
		Role:     reqBody.Role,
	})

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		logrus.WithFields(logrus.Fields{
			"domain":  "User",
			"handler": "Create User",
			"err":     err.Error(),
		}).Error("Create")
		return
	}

	util.SendNoData(w, http.StatusOK, "Success Create User")
	logrus.WithFields(logrus.Fields{
		"domain":   "User",
		"handler":  "Create User",
		"username": reqBody.Username,
	}).Info("Success")
	return
}

// register for user
func (h UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {

	var reqBody RequestRegisterUser

	err := util.Decode(r, &reqBody)

	if err != nil {
		util.SendNoData(w, http.StatusBadRequest, err.Error())
		logrus.WithFields(logrus.Fields{
			"domain":  "User",
			"handler": "Register User",
			"err":     err.Error(),
		}).Error("Decode")
		return
	}

	var validationErrPayload ValidationRegisterUserErrorResponse
	err = h.validator.Validate(reqBody, &validationErrPayload)

	if err != nil {
		util.SendWithData(w, http.StatusUnprocessableEntity, err.Error(), validationErrPayload)
		logrus.WithFields(logrus.Fields{
			"domain":  "User",
			"handler": "Register User",
			"err":     err.Error(),
		}).Error("Validate")
		return
	}

	err = h.usService.Create(r.Context(), entities.User{
		Username: reqBody.Username,
		Email:    reqBody.Email,
		Password: reqBody.Password,
		Fullname: reqBody.Fullname,
		Role:     2,
	})

	if err != nil {
		errors, ok := err.(*helper.Errors)
		if ok {
			util.SendWithData(w, http.StatusUnprocessableEntity, errors.Error(), errors.Data)
			logrus.WithFields(logrus.Fields{
				"domain":  "User",
				"handler": "Register User",
				"err":     err.Error(),
			}).Error("Create")
			return
		}
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		logrus.WithFields(logrus.Fields{
			"domain":  "User",
			"handler": "Register User",
			"err":     err.Error(),
		}).Error("Create")
		return
	}

	util.SendNoData(w, http.StatusOK, "Success Create User")
	logrus.WithFields(logrus.Fields{
		"domain":   "User",
		"handler":  "Register User",
		"username": reqBody.Username,
	}).Info("Success")
	return
}
