package transaction

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
	"github.com/sirupsen/logrus"
)

func (th TransactionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	tID := util.GetParams(r, "id")

	// call get by id service
	t, err := th.service.GetByID(r.Context(), tID)
	if err != nil {
		util.SendNoData(w, http.StatusNotFound, "No transaction found")

		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Get By Id",
			"err":     err.Error(),
		}).Error("Get By Id")

		return
	}

	util.SendWithData(w, http.StatusOK, "", ResponseBody{
		ID:              t.ID,
		TransactionName: t.TransactionName,
		Category:        t.Category,
		TransactionType: t.TransactionType,
		Amount:          t.Amount,
		TransactionAt:   t.TransactionAt,
		CreatedAt:       t.CreatedAt,
		UpdatedAt:       t.UpdatedAt,
		DeletedAt:       t.DeletedAt,
	})

	logrus.WithFields(logrus.Fields{
		"domain":         "Transaction",
		"handler":        "Get By Id",
		"transaction_id": tID,
	}).Error("Success")

	return
}
