package repo

import (
	"github.com/adityjoshi/E-Commerce-/service/authService/db"
	"github.com/adityjoshi/E-Commerce-/service/authService/models"
)

type UserRepo struct{}

func (u *UserRepo) CreateUser(user *models.Users) error {
	database := db.GetDB(user.Region)
	result := database.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
