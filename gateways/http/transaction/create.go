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
		// _ = util.SendError(w, err.Error(), http.StatusBadRequest, nil)
		util.SendNoData(w, http.StatusBadRequest, err.Error())
		return
	}

	var reqValidation TransactionRequestValidationError
	err = th.validator.Validate(reqBody, reqValidation)
	if err != nil {
		// _ = util.SendError(w, util.ErrValidation, http.StatusUnprocessableEntity, reqValidation)
		util.SendWithData(w, http.StatusUnprocessableEntity, util.ErrValidation, reqValidation)
		return
	}

	// #1 - check user_id exist or not
	// #2 - dont have to check userID, the reason why is because if user can access this transaction it means
	// the user exist
	// =================
	// userExist, err := th.service.CheckUser(r.Context(), r.Context().Value("user_id").(int))
	// if err != nil {
	// 	_ = util.SendError(w, err.Error(), http.StatusInternalServerError, nil)
	// 	return
	// }
	// if !userExist {
	// 	_ = util.SendError(w, util.ErrUserNotFound, http.StatusNotFound, nil)
	// 	return
	// }

	err = th.service.Create(r.Context(), entities.TransactionInput{
		TransactionName: reqBody.TransactionName,
		Category:        reqBody.Category,
		TransactionType: reqBody.TransactionType,
		Amount:          reqBody.Amount,
		TransactionAt:   reqBody.TransactionAt,
	})
	if err != nil {
		// _ = util.SendError(w, err.Error(), http.StatusInternalServerError, nil)
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	// _ = util.SendSuccess(w, "Transaction created successfully!", http.StatusOK, nil)
	util.SendNoData(w, http.StatusOK, "Transaction created successfully!")
}
