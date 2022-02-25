package transaction

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
)

func (th TransactionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	tID := util.GetParams(r, "id")

	// validate transaction_id
	te, err := th.service.CheckTransactionByID(r.Context(), tID)
	if err != nil {
		_ = util.SendError(w, err.Error(), http.StatusNotFound, nil)
		return
	}
	if !te {
		_ = util.SendError(w, "transaction not found", http.StatusNotFound, nil)
		return
	}

	// call delete service
	err = th.service.Delete(r.Context(), tID)
	if err != nil {
		_ = util.SendError(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}

	_ = util.SendSuccess(w, "transaction deleted successfully!", http.StatusOK, nil)
}
