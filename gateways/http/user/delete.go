package user

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
	"github.com/sirupsen/logrus"
)

func (h UserHandler) Delete(w http.ResponseWriter, r *http.Request) {

	userId := util.GetParams(r, "id")

	err := h.usService.Delete(r.Context(), userId)

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		logrus.WithFields(logrus.Fields{
			"domain":  "User",
			"service": "Delete",
			"message": err.Error(),
		}).Error("Error")
		return
	}

	util.SendNoData(w, http.StatusOK, "Success Delete User")
	logrus.WithFields(logrus.Fields{
		"domain":  "User",
		"service": "Delete",
		"user_id": userId,
	}).Info("Success")
	return
}
