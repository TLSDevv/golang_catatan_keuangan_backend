package transaction

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
	"github.com/sirupsen/logrus"
)

func (th TransactionHandler) Update(w http.ResponseWriter, r *http.Request) {
	tID := util.GetParams(r, "id")
	var reqBody entities.TransactionInput

	err := util.Decode(r, &reqBody)
	if err != nil {
		_ = util.SendError(w, err.Error(), http.StatusBadRequest, nil)
		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Update",
			"err":     err.Error(),
		}).Error("Decode")
		return
	}

	errs := reqBody.Validate()
	if errs != nil {
		_ = util.SendError(w, util.ErrValidation, http.StatusUnprocessableEntity, errs)
		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Update",
			"err":     err.Error(),
		}).Error("Validate")
		return
	}

	// validate user_id
	ue, err := th.service.CheckUser(r.Context(), r.Context().Value(util.CtxUserId).(int))
	if err != nil {
		_ = util.SendError(w, err.Error(), http.StatusInternalServerError, nil)
		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Update",
			"err":     err.Error(),
		}).Error("Check User")
		return
	}
	if !ue {
		_ = util.SendError(w, util.ErrUserNotFound, http.StatusNotFound, nil)
		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Update",
			"err":     err.Error(),
		}).Error("User Not Found")
		return
	}

	// validate transaction_id
	te, err := th.service.CheckTransactionByID(r.Context(), tID)
	if err != nil {
		_ = util.SendError(w, err.Error(), http.StatusNotFound, nil)
		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Update",
			"err":     err.Error(),
		}).Error("Check Transaction By Id")
		return
	}
	if !te {
		_ = util.SendError(w, "transaction not found", http.StatusNotFound, nil)
		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Update",
			"err":     err.Error(),
		}).Error("Transaction Not Found")
		return
	}

	// call update service
	err = th.service.Update(r.Context(), reqBody, tID)
	if err != nil {
		_ = util.SendError(w, err.Error(), http.StatusInternalServerError, nil)
		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Update",
			"err":     err.Error(),
		}).Error("Update")
		return
	}

	_ = util.SendSuccess(w, "transaction updated successfully!", http.StatusOK, nil)
	logrus.WithFields(logrus.Fields{
		"domain":         "Transaction",
		"handler":        "Update",
		"user_id":        r.Context().Value(util.CtxUserId),
		"transaction_id": tID,
	}).Info("Success")
}
