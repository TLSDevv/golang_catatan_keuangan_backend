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
				util.SendNoData(w, http.StatusBadRequest, err.Error())
				return
			}

			ctx := context.WithValue(context.Background(), "user_id", accessDetails.Id)
			ctx = context.WithValue(ctx, "username", accessDetails.Username)
			r.WithContext(ctx)

			h.ServeHTTP(w, r)
		})
}
