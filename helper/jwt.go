package helper

import (
	"f-commerce/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type Jwt struct {
	Cfg *config.Config
	Log *zap.Logger
}

type CustomClaim struct {
	ID    string
	Email string
	Role  string
	jwt.RegisteredClaims
}

func NewJwt(Cfg *config.Config, Log *zap.Logger) *Jwt {
	return &Jwt{Cfg, Log}
}

func (j *Jwt) CreateToken(email, id, role string) (string, error) {

	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(j.Cfg.Key.PrivateKey))
	if err != nil {
		j.Log.Error("failed to parse private key", zap.Error(err))
		return "", err
	}

	expiration := time.Now().Add(24 * time.Hour)

	claims := &CustomClaim{
		ID:    id,
		Email: email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			Issuer:    j.Cfg.AppName,
			Subject:   id,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
