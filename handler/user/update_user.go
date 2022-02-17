package user

import (
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/handler/util"
)

func (h UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	var reqBody RequestUpdateUser

	err := util.Decode(r, &reqBody)

	userId := util.GetParams(r, "id")

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	var validationErrPayload ValidationErrorResponse
	err = h.validator.Validate(reqBody, &validationErrPayload)

	if err != nil {
		util.SendNoData(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.usService.UpdateUser(r.Context(), userId, entities.User{
		Username: reqBody.Username,
		Email:    reqBody.Email,
		Fullname: reqBody.Fullname,
	})

	if err != nil {
		util.SendNoData(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendNoData(w, http.StatusOK, "Success Update User")
	return
}
