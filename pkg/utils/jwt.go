package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("MyVeryVerySecretKey")

type JWTClaims struct {
	UserId   int    `json:"user_id"`
	UserType string `json:"user_type"`
	jwt.StandardClaims
}

// JWT oluşturma fonksiyonu
func GenerateJWT(userId int, userType string) (string, error) {
	claims := &JWTClaims{
		UserId:   userId,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 24 saat geçerli olacak
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// JWT doğrulama fonksiyonu
func ValidateJWT(jwtToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(
		jwtToken,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New("Invalid JWT claims")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token expired")
	}

	return claims, nil
}
