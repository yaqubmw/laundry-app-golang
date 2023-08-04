package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     []byte
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

type TokenMyClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}
