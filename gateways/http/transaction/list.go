package transaction

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
	"github.com/sirupsen/logrus"
)

func (th TransactionHandler) List(w http.ResponseWriter, r *http.Request) {
	// call get transactions service
	transactions, err := th.service.GetTransactions(r.Context())
	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())

		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "List",
			"err":     err.Error(),
		}).Error("Get Transactions")

		return
	}

	transactionsResponse := formatSliceResponse(transactions)
	util.SendWithData(w, http.StatusOK, "", transactionsResponse)

	logrus.WithFields(logrus.Fields{
		"domain":            "Transaction",
		"handler":           "List",
		"total_transaction": len(transactionsResponse),
	}).Info("Success")

	return
}
