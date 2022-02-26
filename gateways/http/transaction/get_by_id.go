package transaction

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
	"github.com/sirupsen/logrus"
)

func (th TransactionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	tID := util.GetParams(r, "id")
	t, err := th.service.GetByID(r.Context(), tID)
	if err != nil {
		_ = util.SendError(w, "No transaction found", http.StatusNotFound, nil)
		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Get By Id",
			"err":     err.Error(),
		}).Error("Get By Id")
		return
	}

	_ = util.SendSuccess(w, "", http.StatusOK, ResponseBody{
		ID:              t.ID,
		TransactionName: t.TransactionName,
		Category:        t.Category,
		TransactionType: t.TransactionType,
		TransactionAt:   t.TransactionAt,
		CreatedAt:       t.CreatedAt,
	})
	logrus.WithFields(logrus.Fields{
		"domain":         "Transaction",
		"handler":        "Get By Id",
		"transaction_id": tID,
	}).Error("Success")
	return
}
