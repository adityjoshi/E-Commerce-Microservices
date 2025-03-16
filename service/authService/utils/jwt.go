package utils

import (
	"fmt"
	"time"

	"github.com/adityjoshi/E-Commerce-/service/authService/models"
	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("JWT123")

func GetJWTKey(user *models.Users) (string, error) {
	claims := jwt.MapClaims{
		"ID":        user.ID,
		"Email":     user.Email,
		"User_type": user.User_type,
		"region":    user.Region,
		"exp":       time.Now().Add(time.Minute * 2).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("could not sign token: %v", err)
	}
	return tokenString, nil
}
