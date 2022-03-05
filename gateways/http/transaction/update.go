package transaction

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
)

func (th TransactionHandler) Update(w http.ResponseWriter, r *http.Request) {
	tID := util.GetParams(r, "id")
	var reqBody TransactionRequest

	err := util.Decode(r, &reqBody)
	if err != nil {
		util.SendNoData(w, http.StatusBadRequest, err.Error())
		return
	}

	// errs := reqBody.Validate()
	// if errs != nil {
	// 	_ = util.SendError(w, util.ErrValidation, http.StatusUnprocessableEntity, errs)
	// 	return
	// }

	var reqValidation TransactionRequestValidationError
	err = th.validator.Validate(reqBody, reqValidation)
	if err != nil {
		util.SendWithData(w, http.StatusUnprocessableEntity, util.ErrValidation, reqValidation)
		return
	}

	/* validate user_id, no need to check user_id, we assume that the user_id is already exist */
	// ue, err := th.service.CheckUser(r.Context(), r.Context().Value("user_id").(int))
	// if err != nil {
	// 	_ = util.SendError(w, err.Error(), http.StatusInternalServerError, nil)
	// 	return
	// }
	// if !ue {
	// 	_ = util.SendError(w, util.ErrUserNotFound, http.StatusNotFound, nil)
	// 	return
	// }

	/* validate transaction_id, no need to check the transaction_id, we assume
	that the transaction_id already exist on the db */
	// te, err := th.service.CheckTransactionByID(r.Context(), tID)
	// if err != nil {
	// 	_ = util.SendError(w, err.Error(), http.StatusNotFound, nil)
	// 	return
	// }
	// if !te {
	// 	_ = util.SendError(w, "transaction not found", http.StatusNotFound, nil)
	// 	return
	// }

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
		return
	}

	util.SendNoData(w, http.StatusOK, "transaction updated successfully!")
	return
}
