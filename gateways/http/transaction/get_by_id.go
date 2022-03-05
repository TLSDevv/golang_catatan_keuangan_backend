package transaction

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
)

func (th TransactionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	tID := util.GetParams(r, "id")
	t, err := th.service.GetByID(r.Context(), tID)
	if err != nil {
		util.SendNoData(w, http.StatusNotFound, "No transaction found")
		return
	}

	util.SendWithData(w, http.StatusOK, "", ResponseBody{
		ID:              t.ID,
		TransactionName: t.TransactionName,
		Category:        t.Category,
		TransactionType: t.TransactionType,
		TransactionAt:   t.TransactionAt,
		CreatedAt:       t.CreatedAt,
	})
	return
}
