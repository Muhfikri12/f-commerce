package authservice

import (
	"errors"
	"f-commerce/helper"
	"f-commerce/model"
	"f-commerce/repository"
	"fmt"
	"strconv"

	"go.uber.org/zap"
)

type AuthService interface {
	Login(login *model.Login) (*model.User, string, error)
	VerificationEmail(verify *model.VerificationEmail) error
	AskNewOTP(email string) error
}

type authService struct {
	repo *repository.Repository
	log  *zap.Logger
	jwt  *helper.Jwt
}

func NewAuthService(repo *repository.Repository, log *zap.Logger, jwt *helper.Jwt) AuthService {
	return &authService{repo, log, jwt}
}

func (as *authService) Login(login *model.Login) (*model.User, string, error) {

	user, err := as.repo.User.GetUser(login)
	if err != nil {
		as.log.Error("failed to fatch repository: ", zap.Error(err))
		return nil, "", err
	}

	if user.Status == "unverified" {
		return nil, "", fmt.Errorf("account is unverified")
	}

	if !helper.CheckHashPassword(login.Password, user.Password) {
		as.log.Error("invalid password")
		return nil, "", errors.New("invalid password")
	}

	id := strconv.Itoa(user.Id)

	token, err := as.jwt.CreateToken(user.Email, id, user.Role)
	if err != nil {
		as.log.Error("failed create token: ", zap.Error(err))
		return nil, "", fmt.Errorf("failed create token: " + err.Error())
	}

	return user, token, nil
}

func (as *authService) VerificationEmail(verify *model.VerificationEmail) error {

	if err := as.repo.Auth.VerificationEmail(verify); err != nil {
		return err
	}

	return nil
}

func (as *authService) AskNewOTP(email string) error {

	user := model.Login{
		Input:    email,
		Password: "k",
	}

	if _, err := as.repo.User.GetUser(&user); err != nil {
		return err
	}

	return nil
}

func (as *authService) Logout(token string) error {

	return nil
}
