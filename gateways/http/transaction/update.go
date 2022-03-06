package transaction

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
	"github.com/sirupsen/logrus"
)

func (th TransactionHandler) Update(w http.ResponseWriter, r *http.Request) {
	tID := util.GetParams(r, "id")

	// decode request
	var reqBody TransactionRequest
	err := util.Decode(r, &reqBody)
	if err != nil {
		util.SendNoData(w, http.StatusBadRequest, err.Error())

		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Update",
			"err":     err.Error(),
		}).Error("Decode")

		return
	}

	// validate request
	var reqValidation TransactionRequestValidationError
	err = th.validator.Validate(reqBody, reqValidation)
	if err != nil {
		util.SendWithData(w, http.StatusUnprocessableEntity, util.ErrValidation, reqValidation)

		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Update",
			"err":     err.Error(),
		}).Error("Validate")

		return
	}

	// call update service
	err = th.service.Update(r.Context(), entities.TransactionInput{
		TransactionName: reqBody.TransactionName,
		Category:        reqBody.Category,
		TransactionType: reqBody.TransactionType,
		Amount:          reqBody.Amount,
		TransactionAt:   reqBody.TransactionAt,
	}, tID)
	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())

		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Update",
			"err":     err.Error(),
		}).Error("Update")

		return
	}

	util.SendNoData(w, http.StatusOK, "Transaction updated successfully!")

	logrus.WithFields(logrus.Fields{
		"domain":         "Transaction",
		"handler":        "Update",
		"user_id":        r.Context().Value(util.CtxUserId),
		"transaction_id": tID,
	}).Info("Success")

	return
}
