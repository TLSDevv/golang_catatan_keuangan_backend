package transaction

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
)

func (th TransactionHandler) Restore(w http.ResponseWriter, r *http.Request) {
	tID := util.GetParams(r, "id")

	// call restore service
	err := th.service.Restore(r.Context(), tID)
	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendNoData(w, http.StatusOK, "Transaction restored successfully!")
	return
}
