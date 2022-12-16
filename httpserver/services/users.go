package services

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/storyofhis/xtrame/common"
	"github.com/storyofhis/xtrame/config"
	"github.com/storyofhis/xtrame/httpserver/controllers/params"
	"github.com/storyofhis/xtrame/httpserver/controllers/views"
	"github.com/storyofhis/xtrame/httpserver/repository"
	"github.com/storyofhis/xtrame/httpserver/repository/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userSvc struct {
	repos repository.UserRepo
}

func NewUserSvc(repos repository.UserRepo) UserSvc {
	return &userSvc{
		repos: repos,
	}
}

func (svc *userSvc) Register(ctx context.Context, user *params.Register) *views.Response {
	_, err := svc.repos.FindUserByEmail(ctx, user.Email)
	if err == nil {
		return views.ErrorResponse(http.StatusBadRequest, views.M_EMAIL_ALREADY_USED, errors.New("email already used"))
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}
	// generate password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	// request from user
	input := models.Users{
		FullName: user.FullName,
		NickName: user.NickName,
		UserName: user.UserName,
		Email:    user.Email,
		Age:      user.Age,
		Password: string(hashedPassword),
	}
	if err = svc.repos.CreateUser(ctx, &input); err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}
	return views.SuccessResponse(http.StatusCreated, views.M_CREATED, views.Register{
		Id:        input.Id,
		FullName:  input.FullName,
		NickName:  input.NickName,
		UserName:  input.UserName,
		Email:     input.Email,
		Password:  input.Password,
		Age:       input.Age,
		Role:      input.Role,
		CreatedAt: input.CreatedAt,
	})
}

func (svc *userSvc) Login(ctx context.Context, user *params.Login) *views.Response {
	model, err := svc.repos.FindUserByEmail(ctx, user.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_INVALID_CREDENTIALS, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(user.Password))
	if err != nil {
		return views.ErrorResponse(http.StatusBadRequest, views.M_INVALID_CREDENTIALS, err)
	}

	role := string(model.Role)

	claims := &common.CustomClaims{
		Id:   int(model.Id),
		Role: role,
	}

	claims.ExpiresAt = time.Now().Add(time.Minute * time.Duration(config.GetJwtExpiredTime())).Unix()
	claims.Subject = model.Email

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(config.GetJwtSignature())

	return views.SuccessResponse(http.StatusOK, views.M_OK, views.Login{
		Token: ss,
	})
}
