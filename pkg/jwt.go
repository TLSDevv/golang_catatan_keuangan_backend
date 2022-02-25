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
	method           = jwt.SigningMethodHS256
	TypeTokenAccess  = 1
	TypeTokenRefresh = 2
	JWT              *JWTEnv
)

type JWTEnv struct {
	RefreshToken         string
	AccessToken          string
	AccessTokenLifeTime  int64
	RefreshTokenLifeTime int64
}

func init() {
	JWT = load()
}

func load() *JWTEnv {
	accessToken := viper.GetString("JWT_TOKEN")
	if len(accessToken) == 0 {
		accessToken = "secret"
	}

	refreshToken := viper.GetString("JWT_REFRESH_TOKEN")
	if len(refreshToken) == 0 {
		refreshToken = "refreshSecret"
	}

	accessTokenLifeTime := viper.GetInt("JWT_TOKEN_LIFETIME")
	if accessTokenLifeTime == 0 {
		accessTokenLifeTime = 24
	}

	refreshTokenLifeTime := viper.GetInt("JWT_REFRESH_TOKEN_LIFETIME")
	if refreshTokenLifeTime == 0 {
		refreshTokenLifeTime = 48
	}

	return &JWTEnv{
		AccessToken:          accessToken,
		AccessTokenLifeTime:  int64(accessTokenLifeTime),
		RefreshToken:         refreshToken,
		RefreshTokenLifeTime: int64(refreshTokenLifeTime),
	}
}

type TokenClaims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
	jwt.StandardClaims
}

type AccessDetails struct {
	Id       int
	Username string
	Role     int
}

func PrepareAccessTokenClaims(username string, id, role int) TokenClaims {
	return TokenClaims{
		ID:       id,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(JWT.AccessTokenLifeTime)).Unix(),
		},
	}

}

func PrepareRefreshTokenClaims(username string, id int) TokenClaims {

	return TokenClaims{
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(JWT.RefreshTokenLifeTime)).Unix(),
		},
	}

}

func GenerateToken(tokenClaims TokenClaims, typeToken int) (string, error) {
	if typeToken == TypeTokenAccess {
		return signedToken(tokenClaims, JWT.AccessToken)
	}
	return signedToken(tokenClaims, JWT.RefreshToken)
}

func signedToken(claims TokenClaims, signed string) (string, error) {
	at := jwt.NewWithClaims(method, claims)
	token, err := at.SignedString([]byte(signed))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if strArr[0] != "Bearer" {
		return ""
	}

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
			return []byte(JWT.AccessToken), nil
		}

		return []byte(JWT.RefreshToken), nil
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

		id := int(claims["id"].(float64))
		role := int(claims["role"].(float64))

		username, ok := claims["username"].(string)
		if !ok {
			return nil, err
		}

		return &AccessDetails{
			Id:       id,
			Username: username,
			Role:     role,
		}, nil
	}
	return nil, err
}

func RefreshTokenIsValid(refreshToken string) bool {
	token, _ := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(JWT.RefreshToken), nil
	})
	// Token Valid Check
	if !token.Valid {
		return false
	}

	return true
}
