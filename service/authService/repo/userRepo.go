package repo

import (
	"fmt"

	"github.com/adityjoshi/E-Commerce-/service/authService/db"
	"github.com/adityjoshi/E-Commerce-/service/authService/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct{}

func (u *UserRepo) CreateUser(user *models.Users) error {
	database := db.GetDB(user.Region)
	if database == nil {
		return fmt.Errorf("invalid region")
	}
	return database.Create(user).Error
}

func (u *UserRepo) UserLogin(email, password, region string) (*models.UserLogin, error) {
	database := db.GetDB(region)
	if database == nil {
		return nil, fmt.Errorf("invalid region")
	}
	var user models.UserLogin
	if err := database.Where("email=?", email).First(&user).Error; err != nil {
		return nil, fmt.Errorf("email not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("invalid password")
	}
	return &user, nil
}
