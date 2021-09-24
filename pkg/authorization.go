package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/helper"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/pkg/constant"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/pkg/jwt"

	"github.com/spf13/viper"
)

type Authenticator struct {
}

func NewAuthenticator() *Authenticator {
	return &Authenticator{}
}

func (authenticator *Authenticator) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		token, ok := authenticator.extractToken(r)
		if !ok {
			helper.PanicIfError(errors.New("Failed Auth"))
		}

		claims, err := jwt.ParseClaims(token)
		if err != nil {
			helper.PanicIfError(err)
		}

		userID := claims.(*jwt.TokenClaims).Id
		if userID == "" {
			helper.PanicIfError(errors.New("Failed User Id"))
		}

		issuer := claims.(*jwt.TokenClaims).Issuer
		if issuer == "" {
			helper.PanicIfError(errors.New("Failed issuer"))
		}

		name := claims.(*jwt.TokenClaims).Name
		username := claims.(*jwt.TokenClaims).Username

		ctx = context.WithValue(ctx, "username", username)
		ctx = context.WithValue(ctx, "name", name)
		ctx = context.WithValue(ctx, "userId", userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (authenticator *Authenticator) extractToken(r *http.Request) (string, bool) {
	token := r.Header.Get("Authorization")
	s := strings.Fields(token)
	if len(s) != 2 {
		return "", false
	}
	if s[0] != "Bearer" {
		return "", false
	}
	return s[1], true
}

func enableCors(w *http.ResponseWriter) {
	origins := viper.GetString("api.allowed_origins")
	if origins == "" {
		origins = strings.Join(constant.DefaultAllowdOrigins, ",")
	}

	methods := viper.GetString("api.allowed_methods")
	if methods == "" {
		methods = strings.Join(constant.DefaultAllowedMethods, ",")
	}

	(*w).Header().Set("Access-Control-Allow-Origin", origins)
	(*w).Header().Set("Access-Control-Allow-Methods", methods)
}
