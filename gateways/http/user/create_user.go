package user

import (
	"fmt"
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
		logrus.Error(err.Error())
		return
	}

	var validationErrPayload ValidationErrorResponse
	err = h.validator.Validate(reqBody, &validationErrPayload)

	if err != nil {
		util.SendWithData(w, http.StatusUnprocessableEntity, err.Error(), validationErrPayload)
		logrus.Error(err.Error())
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
		logrus.Error(err.Error())
		return
	}

	util.SendNoData(w, http.StatusOK, "Success Create User")
	logrus.Info("Success Create User")
	return
}

// register for user
func (h UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {

	var reqBody RequestRegisterUser

	err := util.Decode(r, &reqBody)

	if err != nil {
		util.SendNoData(w, http.StatusBadRequest, err.Error())
		logrus.Error(err.Error())
		return
	}

	var validationErrPayload ValidationRegisterUserErrorResponse
	err = h.validator.Validate(reqBody, &validationErrPayload)

	if err != nil {
		util.SendWithData(w, http.StatusUnprocessableEntity, err.Error(), validationErrPayload)
		logrus.Error(err.Error())
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
		fmt.Println(ok)
		if ok {
			util.SendWithData(w, http.StatusUnprocessableEntity, errors.Error(), errors.Data)
			logrus.Error(err.Error())
			return
		}
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		logrus.Error(err.Error())
		return
	}

	util.SendNoData(w, http.StatusOK, "Success Create User")
	logrus.Info("Success Create User")
	return
}
