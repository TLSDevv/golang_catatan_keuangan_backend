package transaction

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
	"github.com/sirupsen/logrus"
)

func (th TransactionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var reqBody entities.TransactionInput

	err := util.Decode(r, &reqBody)
	if err != nil {
		_ = util.SendError(w, err.Error(), http.StatusBadRequest, nil)
		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Create",
			"err":     err.Error(),
		}).Error("Decode")
		return
	}

	errs := reqBody.Validate()
	if errs != nil {
		_ = util.SendError(w, util.ErrValidation, http.StatusUnprocessableEntity, errs)
		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Create",
			"err":     err.Error(),
		}).Error("Validate")
		return
	}

	// check user_id exist or not
	userExist, err := th.service.CheckUser(r.Context(), r.Context().Value(util.CtxUserId).(int))
	if err != nil {
		_ = util.SendError(w, err.Error(), http.StatusInternalServerError, nil)
		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Create",
			"err":     err.Error(),
		}).Error("Check User")
		return
	}
	if !userExist {
		_ = util.SendError(w, util.ErrUserNotFound, http.StatusNotFound, nil)
		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Create",
			"err":     err.Error(),
		}).Error("Decode")
		return
	}

	err = th.service.Create(r.Context(), reqBody)
	if err != nil {
		_ = util.SendError(w, err.Error(), http.StatusInternalServerError, nil)
		logrus.WithFields(logrus.Fields{
			"domain":  "Transaction",
			"handler": "Create",
			"err":     err.Error(),
		}).Error("Create")
		return
	}

	_ = util.SendSuccess(w, "Transaction created successfully!", http.StatusOK, nil)
	logrus.WithFields(logrus.Fields{
		"domain":      "Transaction",
		"handler":     "Create",
		"transaction": reqBody.TransactionName,
	}).Info("Success")
	return
}
