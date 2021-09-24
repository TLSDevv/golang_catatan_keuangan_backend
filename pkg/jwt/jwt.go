package jwt

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/config"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/pkg/constant"
	jwt "github.com/dgrijalva/jwt-go"
)

type TokenClaims struct {
	*jwt.StandardClaims
	Username string `json:"username"`
	Name     string `json:"name"`
	Id       string `json:"id"`
}

type UserClaims interface {
	GetId() string
	GetName() string
	GetUsername() string
}

var verifyKey *rsa.PublicKey
var signKey *rsa.PrivateKey
var tokenExpiration map[string]time.Duration

func Init(config *config.Config) error {
	initTokenExpiration(config)

	time.Local = constant.JakartaTimezone
	privateKeyPath := config.AUTH.JWTPrivateKey
	pubKeyPath := config.AUTH.JWTPublicKey

	signBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return fmt.Errorf("Can't read private key: %w", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return fmt.Errorf("Can't read private key: %w", err)
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		return fmt.Errorf("Can't read public key: %w", err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return fmt.Errorf("Can't read public key: %w", err)
	}

	return nil
}

func initTokenExpiration(cfg *config.Config) {
	tokenExpiration = cfg.AUTH.JWTLifetime
}

func GenerateJWT(claims *TokenClaims) (string, error) {
	t := jwt.New(jwt.SigningMethodRS256)

	t.Claims = claims
	token, err := t.SignedString(signKey)
	return token, err
}

func ParseClaims(tokenString string) (interface{}, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Now().In(constant.JakartaTimezone)
	}
	tokenString = strings.TrimSpace(tokenString)

	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})

	if err != nil {
		return nil, err
	}
	claims := token.Claims.(*TokenClaims)
	return claims, nil
}

type asd struct {
	asd string
}

func PrepareAccessTokenClaims(tokenLifetimeIdentifier, issuer, sessionID string, userDetail UserClaims) *TokenClaims {
	lifetime := tokenExpiration[tokenLifetimeIdentifier]
	expireTime := time.Now().In(constant.JakartaTimezone).Add(lifetime).Unix()

	return &TokenClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    issuer,
			Id:        sessionID,
		},
		Username: userDetail.GetUsername(),
		Name:     userDetail.GetName(),
		Id:       userDetail.GetId(),
	}
}

func PrepareRefreshTokenClaims(tokenLifetimeIdentifier, issuer, sessionID string, userDetail UserClaims) *TokenClaims {
	lifetime := tokenExpiration[tokenLifetimeIdentifier]
	expireTime := time.Now().In(constant.JakartaTimezone).Add(lifetime).Unix()

	return &TokenClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    issuer,
			Id:        sessionID,
			Subject:   "",
		},
		Username: userDetail.GetUsername(),
		Name:     userDetail.GetName(),
		Id:       userDetail.GetId(),
	}
}
