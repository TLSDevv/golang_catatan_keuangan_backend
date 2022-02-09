package pkg

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var (
	method               = jwt.SigningMethodHS256
	RefreshToken         string
	AccessToken          string
	AccessTokenLifeTime  int
	RefreshTokenLifeTime int
	TypeTokenAccess      = 1
	TypeTokenRefresh     = 2
)

func init() {
	token := viper.GetString("JWT_TOKEN")
	if len(token) == 0 {
		token = "secret"
	}

	refreshToken := viper.GetString("JWT_REFRESH_TOKEN")
	if len(refreshToken) == 0 {
		refreshToken = "refreshSecret"
	}

	tokenLifeTime := viper.GetInt("JWT_TOKEN_LIFETIME")
	if tokenLifeTime == 0 {
		tokenLifeTime = 24
	}

	refreshTokenLifeTime := viper.GetInt("JWT_REFRESH_TOKEN_LIFETIME")
	if refreshTokenLifeTime == 0 {
		refreshTokenLifeTime = 48
	}
}

type TokenClaims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type AccessDetails struct {
	Id       int
	Username string
}

func PrepareTokenClaims(username string, id, typeToken int) *TokenClaims {
	if typeToken == TypeTokenAccess {
		return &TokenClaims{
			ID:       id,
			Username: username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * time.Duration(AccessTokenLifeTime)).Unix(),
			},
		}
	}

	return &TokenClaims{
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(RefreshTokenLifeTime)).Unix(),
		},
	}
}

func GenerateToken(tokenClaims *TokenClaims, typeToken int) (string, error) {
	if typeToken == TypeTokenAccess {
		return signedToken(tokenClaims, AccessToken)
	}
	return signedToken(tokenClaims, RefreshToken)
}

func signedToken(claims *TokenClaims, signed string) (string, error) {
	at := jwt.NewWithClaims(method, claims)
	token, err := at.SignedString(signed)

	if err != nil {
		return "", err
	}

	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request, typeToken int) (*jwt.Token, error) {
	tokenString := ExtractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		if typeToken == TypeTokenAccess {
			return AccessToken, nil
		}

		return RefreshToken, nil
	})

	if err != nil {
		return nil, err
	}

	// Token Valid Check
	if _, ok := token.Claims.(TokenClaims); !ok && !token.Valid {
		return nil, err
	}

	return token, nil
}

func ExtractTokenMetadata(r *http.Request, typeToken int) (*AccessDetails, error) {
	token, err := VerifyToken(r, typeToken)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		id, ok := claims["id"].(int)
		if !ok {
			return nil, err
		}
		username, ok := claims["username"].(string)
		if !ok {
			return nil, err
		}
		return &AccessDetails{
			Id:       id,
			Username: username,
		}, nil
	}
	return nil, err
}
