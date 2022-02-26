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
				// [jamil] - change to unauthorized 401
				util.SendNoData(w, http.StatusUnauthorized, err.Error())
				return
			}

			// [jamil] - change context.Background() to r.Context()
			ctx := context.WithValue(r.Context(), "user_id", accessDetails.Id)
			ctx = context.WithValue(ctx, "username", accessDetails.Username)
			ctx = context.WithValue(ctx, "role", accessDetails.Role)

			// [jamil] - redefine the request
			r = r.WithContext(ctx)

			h.ServeHTTP(w, r)
		})
}

func Admin(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			role := r.Context().Value("role")

			if role != 1 {
				util.SendNoData(w, http.StatusForbidden, "Forbidden Access")
				return
			}

			h.ServeHTTP(w, r)
		})
}
