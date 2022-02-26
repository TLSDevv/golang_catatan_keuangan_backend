package user

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
	"github.com/sirupsen/logrus"
)

func (h UserHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {

	var reqBody RequestUpdateRole

	err := util.Decode(r, &reqBody)

	userId := util.GetParams(r, "id")

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		logrus.WithFields(logrus.Fields{
			"domain":  "User",
			"handler": "Update_Role",
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
			"handler": "Update Role",
			"err":     err.Error(),
		}).Error("Validate")
		return
	}

	err = h.usService.UpdateRole(r.Context(), userId, reqBody.Role)

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		logrus.WithFields(logrus.Fields{
			"domain":  "User",
			"handler": "Update Role",
			"err":     err.Error(),
		}).Error("Update Role")
		return
	}

	util.SendNoData(w, http.StatusOK, "Success Update User Role")
	logrus.Error("Success Update User Role")
	logrus.WithFields(logrus.Fields{
		"domain":  "User",
		"handler": "Update Role",
		"user_id": userId,
	}).Info("Success")
	return
}
