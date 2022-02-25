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
		logrus.Error(err.Error())
		return
	}

	var validationErrPayload ValidationErrorResponse
	err = h.validator.Validate(reqBody, &validationErrPayload)

	if err != nil {
		util.SendNoData(w, http.StatusBadRequest, err.Error())
		logrus.Error(err.Error())
		return
	}

	err = h.usService.UpdateRole(r.Context(), userId, reqBody.Role)

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		logrus.Error(err.Error())
		return
	}

	util.SendNoData(w, http.StatusOK, "Success Update User Role")
	logrus.Error("Success Update User Role")
	return
}
