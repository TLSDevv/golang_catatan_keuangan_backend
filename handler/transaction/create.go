package transaction

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/handler/util"
)

func (th TransactionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var reqBody entities.CreateTransactionInput

	err := util.Decode(r, &reqBody)
	if err != nil {
		_ = util.SendError(w, err.Error(), http.StatusBadRequest, nil)
		return
	}

	errs := reqBody.Validate()
	if errs != nil {
		_ = util.SendError(w, util.ErrValidation, http.StatusUnprocessableEntity, errs)
		return
	}

	// validate user_id
	userExist, err := th.service.CheckUser(r.Context(), reqBody.UserID)
	if err != nil {
		_ = util.SendError(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}
	if !userExist {
		_ = util.SendError(w, util.ErrUserNotFound, http.StatusNotFound, nil)
		return
	}

	err = th.service.Create(r.Context(), reqBody)
	if err != nil {
		_ = util.SendError(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}

	_ = util.SendSuccess(w, "Transaction created successfully!", http.StatusOK, nil)
}
