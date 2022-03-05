package transaction

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
)

func (th TransactionHandler) List(w http.ResponseWriter, r *http.Request) {
	transactions, err := th.service.GetTransactions(r.Context())
	if err != nil {
		// _ = util.SendError(w, err.Error(), http.StatusInternalServerError, nil)
		// return
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	transactionsResponse := formatSliceResponse(transactions)
	// _ = util.Send(w, TransactionListResponse{Success: true, Data: transactionsResponse}, http.StatusOK)
	util.SendWithData(w, http.StatusOK, "", transactionsResponse)
	return
}
