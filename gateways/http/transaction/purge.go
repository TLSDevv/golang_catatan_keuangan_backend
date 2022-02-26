package transaction

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
	"github.com/sirupsen/logrus"
)

func (th TransactionHandler) Purge(w http.ResponseWriter, r *http.Request) {
	tID := util.GetParams(r, "id")
	// validate transaction_id
	te, err := th.service.CheckTransactionByID(r.Context(), tID)
	if err != nil {
		_ = util.SendError(w, err.Error(), http.StatusNotFound, nil)
		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Purge",
			"err":     err.Error(),
		}).Error("Check Transaction By Id")
		return
	}
	if !te {
		_ = util.SendError(w, "transaction not found", http.StatusNotFound, nil)
		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Purge",
			"err":     err.Error(),
		}).Error("Transaction Not Found")
		return
	}

	// call purge service
	err = th.service.Purge(r.Context(), tID)
	if err != nil {
		_ = util.SendError(w, err.Error(), http.StatusInternalServerError, nil)
		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Purge",
			"err":     err.Error(),
		}).Error("Purge")
		return
	}

	_ = util.SendSuccess(w, "transaction purged successfully!", http.StatusOK, nil)
	logrus.WithFields(logrus.Fields{
		"domain":         "Transaction",
		"handler":        "Purge",
		"transaction_id": tID,
	}).Info("Success")
	return
}
