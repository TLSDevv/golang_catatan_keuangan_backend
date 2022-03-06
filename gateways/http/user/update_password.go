package user

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
	"github.com/sirupsen/logrus"
)

func (h UserHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {

	var reqBody RequestUpdatePassword

	err := util.Decode(r, &reqBody)

	userId := util.GetParams(r, "id")

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		logrus.WithFields(logrus.Fields{
			"domain":  "User",
			"handler": "Update_Password",
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
			"handler": "Update_Password",
			"err":     err.Error(),
		}).Error("Validate")
		return
	}

	err = h.usService.UpdatePassword(r.Context(), userId, reqBody.Password)

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		logrus.WithFields(logrus.Fields{
			"domain":  "User",
			"handler": "Update Password",
			"err":     err.Error(),
		}).Error("Service.UpdatePassword")
		return
	}

	util.SendNoData(w, http.StatusOK, "User updated successfully!")
	logrus.WithFields(logrus.Fields{
		"domain":  "User",
		"handler": "Update Password",
		"user_id": userId,
	}).Info("Success")
	return
}
