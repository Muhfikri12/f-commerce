package authservice

import (
	"errors"
	"finance/helper"
	"finance/model"
	"finance/repository"
	"fmt"
	"strconv"

	"go.uber.org/zap"
)

type AuthService interface {
	Login(login *model.Login) (string, error)
}

type authService struct {
	repo *repository.Repository
	log  *zap.Logger
	jwt  *helper.Jwt
}

func NewAuthService(repo *repository.Repository, log *zap.Logger, jwt *helper.Jwt) AuthService {
	return &authService{repo, log, jwt}
}

func (as *authService) Login(login *model.Login) (string, error) {

	user, err := as.repo.Auth.Login(login)
	if err != nil {
		as.log.Error("failed to fatch repository: ", zap.Error(err))
		return "", err
	}

	if !helper.CheckHashPassword(login.Password, user.Password) {
		as.log.Error("invalid password")
		return "", errors.New("invalid password")
	}

	id := strconv.Itoa(user.Id)

	token, err := as.jwt.CreateToken(user.Email, id, user.Role)
	if err != nil {
		as.log.Error("failed create token: ", zap.Error(err))
		return "", fmt.Errorf("failed create token: " + err.Error())
	}

	return token, nil
}
