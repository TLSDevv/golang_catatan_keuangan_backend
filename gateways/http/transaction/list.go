package transaction

import (
	"fmt"
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
)

func (th TransactionHandler) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Context().Value("user_id"))
	transactions, err := th.service.GetTransactions(r.Context())
	if err != nil {
		_ = util.SendError(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}

	transactionsResponse := formatSliceResponse(transactions)
	_ = util.Send(w, TransactionListResponse{Success: true, Data: transactionsResponse}, http.StatusOK)
}
