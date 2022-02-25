package auth

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
	"github.com/sirupsen/logrus"
)

func (handler AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var reqBody RequestLogin

	err := util.Decode(r, &reqBody)

	if err != nil {
		util.SendNoData(w, http.StatusBadRequest, err.Error())
		return
	}

	var validationErrPayload ValidationLoginResponse
	err = handler.validator.Validate(reqBody, &validationErrPayload)

	if err != nil {
		util.SendWithData(w, http.StatusUnprocessableEntity, err.Error(), validationErrPayload)
		logrus.Error(err.Error())
		return
	}

	token, err := handler.AuthService.Login(r.Context(), entities.Login{
		Username: reqBody.Username,
		Password: reqBody.Password,
	})

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendWithData(w, http.StatusOK, "Success Login", token)
	return
}
