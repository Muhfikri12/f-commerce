package helper

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"f-commerce/config"
	"fmt"
	"strconv"
	"strings"
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

func (j *Jwt) ParsingPayload(tokenStr string) (int, error) {

	parts := strings.Split(tokenStr, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		j.Log.Error("Invalid Authorization header format")
		return 0, fmt.Errorf("invalid authorization header format")
	}

	tokenStr = parts[1]

	block, _ := pem.Decode([]byte(j.Cfg.Key.PublicKey))
	if block == nil {
		j.Log.Error("PEM decoding failed: block is nil")
		return 0, fmt.Errorf("PEM decoding failed: block is nil")
	}

	if block.Type != "PUBLIC KEY" {
		j.Log.Error("Unexpected PEM block type", zap.String("type", block.Type))
		return 0, fmt.Errorf("unexpected PEM block type: %s", block.Type)
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		j.Log.Error("Failed to parse public key", zap.Error(err))
		return 0, fmt.Errorf("failed to parse public key: %v", err)
	}

	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		j.Log.Error("Public key is not RSA")
		return 0, fmt.Errorf("not an RSA public key")
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			j.Log.Error("Unexpected signing method", zap.Any("alg", token.Header["alg"]))
			return 0, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return rsaPubKey, nil
	})
	if err != nil {
		j.Log.Error("Error parsing token", zap.Error(err))
		return 0, fmt.Errorf("error parsing token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		idStr, exists := claims["ID"]
		if !exists {
			j.Log.Error("ID not found in token claims")
			return 0, fmt.Errorf("id not found in token")
		}

		id, _ := strconv.Atoi(idStr.(string))

		return id, nil
	}

	j.Log.Error("Invalid token")
	return 0, fmt.Errorf("invalid token")
}
