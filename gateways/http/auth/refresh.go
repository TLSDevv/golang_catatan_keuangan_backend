package auth

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/pkg"
)

func (handler AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {

	accessDetails, err := pkg.ExtractTokenMetadata(r, pkg.TypeTokenRefresh)

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := handler.AuthService.Refresh(r.Context(), accessDetails.Id)

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendWithData(w, http.StatusOK, "Success Refresh Token", token)
	return
}
