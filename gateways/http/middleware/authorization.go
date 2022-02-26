package middleware

import (
	"context"
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/gateways/http/util"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/pkg"
)

func Authorization(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			accessDetails, err := pkg.ExtractTokenMetadata(r, pkg.TypeTokenAccess)

			if err != nil {
				util.SendNoData(w, http.StatusUnauthorized, err.Error())
				return
			}

			ctx := context.WithValue(r.Context(), util.CtxUserId, accessDetails.Id)
			ctx = context.WithValue(ctx, util.CtxUsername, accessDetails.Username)
			ctx = context.WithValue(ctx, util.CtxRole, accessDetails.Role)

			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
}

func Admin(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			role := r.Context().Value(util.CtxRole)

			if role != 1 {
				util.SendNoData(w, http.StatusForbidden, "Forbidden Access")
				return
			}

			h.ServeHTTP(w, r)
		})
}
