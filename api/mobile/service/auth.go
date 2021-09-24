package service

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/mobile/parameter"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/repository"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/exception"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/helper"
	pkgjwt "github.com/TLSDevv/golang_catatan_keuangan_backend/pkg/jwt"

	"github.com/google/uuid"
)

var (
	//register repo
	userRepository = repository.NewUserRepository()
)

type AuthService struct {
	db *sql.DB
}

func NewAuthService(Db *sql.DB) *AuthService {
	return &AuthService{
		db: Db,
	}
}

type Auth struct {
	Username string
	Name     string
	Id       string
}

func (a *Auth) GetUsername() string {
	return a.Username
}

func (a *Auth) GetName() string {
	return a.Name
}

func (a *Auth) GetId() string {
	return a.Id
}

func (authS *AuthService) Login(ctx context.Context, request parameter.LoginRequest) (*parameter.LoginResponse, error) {
	tx, err := authS.db.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	issuer := "mobile"
	sessionId := uuid.New().String()

	user, err := userRepo.GetByUsername(ctx, tx, request.Username)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	userProfile := &Auth{
		Username: user.Username,
		Name:     user.Name,
		Id:       strconv.Itoa(int(user.Id)),
	}

	tokenLifeTime := issuer
	tokenClaims := pkgjwt.PrepareAccessTokenClaims(tokenLifeTime, issuer, sessionId, userProfile)
	accesstoken, err := pkgjwt.GenerateJWT(tokenClaims)

	tokenLifeTime = "refreshtoken"
	tfokenClaims := pkgjwt.PrepareRefreshTokenClaims(tokenLifeTime, issuer, sessionId, userProfile)
	refreshToken, err := pkgjwt.GenerateJWT(tfokenClaims)

	return &parameter.LoginResponse{
		TokenData: parameter.TokenData{
			AccessToken:  accesstoken,
			RefreshToken: refreshToken,
		},
		UserData: helper.ToUserResponse(user),
	}, nil
}

func (authS *AuthService) Logout(ctx context.Context, req parameter.LogoutRequest) (*parameter.LogoutResponse, error) {
	return nil, nil
}
