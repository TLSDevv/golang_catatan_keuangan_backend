package transaction

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/handler/util"
)

func (th TransactionHandler) Restore(w http.ResponseWriter, r *http.Request) {
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

	// call restore service
	err = th.service.Restore(r.Context(), tID)
	if err != nil {
		_ = util.SendError(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}

	_ = util.SendSuccess(w, "transaction restored successfully!", http.StatusOK, nil)
}
