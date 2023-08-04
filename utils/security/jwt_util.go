package security

import (
	"enigma-laundry-apps/config"
	"enigma-laundry-apps/model"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(user model.UserCredential) (string, error) {
	cfg, _ := config.NewConfig()
	now := time.Now().UTC()
	end := now.Add(cfg.AccessTokenLifeTime)

	claims := &TokenMyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.ApplicationName,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Username: user.Username,
	}
	token := jwt.NewWithClaims(cfg.JwtSigningMethod, claims)
	ss, err := token.SignedString(cfg.JwtSignatureKey)
	fmt.Printf("%v %v", ss, err)
	if err != nil {
		return "", fmt.Errorf("failed to create access token: %v", err)
	}
	return ss, nil
}

func VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	cfg, _ := config.NewConfig()

	// digunakan untuk mem-Parse token yang dikirimkan dari Client
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

		// pengecekan sebuah method yang digunakan
		// validasi signature seperti (SIGNING METHOD) yang digunakan yaitu HS256
		method, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != cfg.JwtSigningMethod {
			return nil, fmt.Errorf("signing method invalid")
		}

		// kita kembalikan validasi dari konfigurasi yg sudah divalidasi di atas
		return cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return nil, err
	}

	// cek claims yang sudah didaftarkan sebelumnya
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != cfg.ApplicationName {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
