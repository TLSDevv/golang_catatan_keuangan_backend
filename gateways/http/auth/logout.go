package auth

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
)

func (handler AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	err := handler.AuthService.Logout(r.Context())

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendNoData(w, http.StatusOK, "Success Logout")
	return
}
