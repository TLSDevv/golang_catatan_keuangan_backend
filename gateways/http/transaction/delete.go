package transaction

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
)

func (th TransactionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	tID := util.GetParams(r, "id")

	// call delete service
	err := th.service.Delete(r.Context(), tID)
	if err != nil {
		// _ = util.SendError(w, err.Error(), http.StatusInternalServerError, nil)
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	// _ = util.SendSuccess(w, "transaction deleted successfully!", http.StatusOK, nil)
	util.SendNoData(w, http.StatusOK, "Transaction deleted successfully!")
	return
}
