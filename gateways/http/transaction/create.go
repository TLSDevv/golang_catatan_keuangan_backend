package transaction

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
)

func (th TransactionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var reqBody TransactionRequest

	err := util.Decode(r, &reqBody)
	if err != nil {
		util.SendNoData(w, http.StatusBadRequest, err.Error())
		return
	}

	var reqValidation TransactionRequestValidationError
	err = th.validator.Validate(reqBody, reqValidation)
	if err != nil {
		util.SendWithData(w, http.StatusUnprocessableEntity, util.ErrValidation, reqValidation)
		return
	}

	err = th.service.Create(r.Context(), entities.TransactionInput{
		TransactionName: reqBody.TransactionName,
		Category:        reqBody.Category,
		TransactionType: reqBody.TransactionType,
		Amount:          reqBody.Amount,
		TransactionAt:   reqBody.TransactionAt,
	})
	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendNoData(w, http.StatusOK, "Transaction created successfully!")
}
