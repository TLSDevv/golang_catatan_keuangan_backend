package transaction

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
	"github.com/sirupsen/logrus"
)

func (th TransactionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	tID := util.GetParams(r, "id")

	// validate transaction_id
	te, err := th.service.CheckTransactionByID(r.Context(), tID)
	if err != nil {
		_ = util.SendError(w, err.Error(), http.StatusNotFound, nil)
		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Delete",
			"err":     err.Error(),
		}).Error("Check Transaction By Id")
		return
	}
	if !te {
		_ = util.SendError(w, "transaction not found", http.StatusNotFound, nil)
		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Delete",
			"err":     err.Error(),
		}).Error("Transaction Not Found")
		return
	}

	// call delete service
	err = th.service.Delete(r.Context(), tID)
	if err != nil {
		_ = util.SendError(w, err.Error(), http.StatusInternalServerError, nil)
		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Delete",
			"err":     err.Error(),
		}).Error("Delete")
		return
	}

	_ = util.SendSuccess(w, "transaction deleted successfully!", http.StatusOK, nil)
	logrus.WithFields(logrus.Fields{
		"domain":         "Transaction",
		"handler":        "Delete",
		"transaction_id": tID,
	}).Info("Success")
	return
}
