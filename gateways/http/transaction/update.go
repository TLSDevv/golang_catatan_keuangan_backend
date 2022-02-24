package transaction

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
)

func (th TransactionHandler) Update(w http.ResponseWriter, r *http.Request) {
	tID := util.GetParams(r, "id")
	var reqBody entities.TransactionInput

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
	ue, err := th.service.CheckUser(r.Context(), reqBody.UserID)
	if err != nil {
		_ = util.SendError(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}
	if !ue {
		_ = util.SendError(w, util.ErrUserNotFound, http.StatusNotFound, nil)
		return
	}

	// validate transaction_id
	te, err := th.service.CheckTransactionByID(r.Context(), tID)
	if err != nil {
		_ = util.SendError(w, err.Error(), http.StatusNotFound, nil)
		return
	}
	if !te {
		_ = util.SendError(w, "transaction not found", http.StatusNotFound, nil)
		return
	}

	// call update service
	err = th.service.Update(r.Context(), reqBody)
	if err != nil {
		_ = util.SendError(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}

	_ = util.SendSuccess(w, "transaction updated successfully!", http.StatusOK, nil)
}
