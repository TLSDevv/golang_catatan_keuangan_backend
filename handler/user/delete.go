package user

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/handler/util"
	"github.com/sirupsen/logrus"
)

func (h UserHandler) Delete(w http.ResponseWriter, r *http.Request) {

	userId := util.GetParams(r, "id")

	err := h.usService.Delete(r.Context(), userId)

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		logrus.Error(err.Error())
		return
	}

	util.SendNoData(w, http.StatusOK, "Success Delete User")
	return
}
