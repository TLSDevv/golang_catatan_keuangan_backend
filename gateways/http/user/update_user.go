package user

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
	"github.com/sirupsen/logrus"
)

func (h UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	var reqBody RequestUpdateUser

	err := util.Decode(r, &reqBody)

	userId := util.GetParams(r, "id")

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		logrus.WithFields(logrus.Fields{
			"domain":  "User",
			"handler": "Update User",
			"err":     err.Error(),
		}).Error("Get Params")
		return
	}

	var validationErrPayload ValidationErrorResponse
	err = h.validator.Validate(reqBody, &validationErrPayload)

	if err != nil {
		util.SendNoData(w, http.StatusBadRequest, err.Error())
		logrus.WithFields(logrus.Fields{
			"domain":  "User",
			"handler": "Update User",
			"err":     err.Error(),
		}).Error("Validate")
		return
	}

	err = h.usService.UpdateUser(r.Context(), userId, entities.User{
		Username: reqBody.Username,
		Email:    reqBody.Email,
		Fullname: reqBody.Fullname,
	})

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		logrus.WithFields(logrus.Fields{
			"domain":  "User",
			"handler": "Update User",
			"err":     err.Error(),
		}).Error("Update User")
		return
	}

	util.SendNoData(w, http.StatusOK, "Success Update User")
	logrus.WithFields(logrus.Fields{
		"domain":  "User",
		"handler": "Update User",
		"user_id": userId,
	}).Error("Success")
	return
}
