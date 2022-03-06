package transaction

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
	"github.com/sirupsen/logrus"
)

func (th TransactionHandler) Purge(w http.ResponseWriter, r *http.Request) {
	tID := util.GetParams(r, "id")

	// call purge service
	err := th.service.Purge(r.Context(), tID)
	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())

		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Purge",
			"err":     err.Error(),
		}).Error("Purge")

		return
	}

	util.SendNoData(w, http.StatusOK, "Transaction purged successfully!")

	logrus.WithFields(logrus.Fields{
		"domain":         "Transaction",
		"handler":        "Purge",
		"transaction_id": tID,
	}).Info("Success")

	return
}
