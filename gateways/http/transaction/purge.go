package transaction

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
)

func (th TransactionHandler) Purge(w http.ResponseWriter, r *http.Request) {
	tID := util.GetParams(r, "id")
	// validate transaction_id
	// te, err := th.service.CheckTransactionByID(r.Context(), tID)
	// if err != nil {
	// 	_ = util.SendError(w, err.Error(), http.StatusNotFound, nil)
	// 	return
	// }
	// if !te {
	// 	_ = util.SendError(w, "transaction not found", http.StatusNotFound, nil)
	// 	return
	// }

	// call purge service
	err := th.service.Purge(r.Context(), tID)
	if err != nil {
		// _ = util.SendError(w, err.Error(), http.StatusInternalServerError, nil)
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	// _ = util.SendSuccess(w, "transaction purged successfully!", http.StatusOK, nil)
	util.SendNoData(w, http.StatusOK, "Transaction purged successfully!")
	return
}
