package helper

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"f-commerce/config"
	"fmt"
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

func (j *Jwt) ParsingPayload(tokenStr string) (any, error) {

	block, _ := pem.Decode([]byte(j.Cfg.Key.PublicKey))
	if block == nil {
		j.Log.Error("PEM decoding failed: block is nil")
		return nil, fmt.Errorf("PEM decoding failed: block is nil")
	}

	if block.Type != "PUBLIC KEY" {
		j.Log.Error("unexpected PEM block type: %s", zap.Any("type", block.Type))
		return nil, fmt.Errorf("unexpected PEM block type: %s", block.Type)
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		j.Log.Error("PEM decoding failed: block is nil")
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}

	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		j.Log.Error("PEM decoding failed: block is nil")
		return nil, fmt.Errorf("not an RSA public key")
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			j.Log.Error("PEM decoding failed: block is nil")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return rsaPubKey, nil
	})

	if err != nil {
		j.Log.Error("PEM decoding failed: block is nil")
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, exists := claims["id"]
		if !exists {
			j.Log.Error("PEM decoding failed: block is nil")
			return nil, fmt.Errorf("id not found in token")
		}
		return id, nil
	}

	return nil, fmt.Errorf("invalid token")

}
