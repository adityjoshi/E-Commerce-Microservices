package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWTSECRET"))

func GenerateJWT(userID uint, email string, User_type string, region string) (string, error) {
	claims := jwt.MapClaims{
		"ID":        userID,
		"Email":     email,
		"User_type": User_type,
		"region":    region,
		"exp":       time.Now().Add(time.Minute * 2).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("could not sign token: %v", err)
	}
	return tokenString, nil
}

func DecodeJwt(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
