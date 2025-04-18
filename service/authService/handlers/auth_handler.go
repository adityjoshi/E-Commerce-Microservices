package handlers

import (
	"net/http"

	"github.com/adityjoshi/E-Commerce-Microservices/service/authService/db"
	"github.com/adityjoshi/E-Commerce-Microservices/service/authService/models"
	"github.com/adityjoshi/E-Commerce-Microservices/service/authService/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(c *gin.Context) {
	var registerRequest struct {
		Full_Name     string `json:"Full_Name"`
		GenderInfo    string `json:"GenderInfo"`
		ContactNumber string `json:"ContactNumber"`
		Email         string `json:"Email"`
		Password      string `json:"Password"`
		Region        string `json:"region"`
	}

	if err := c.BindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database := db.GetDB(registerRequest.Region)
	if database == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid region"})
		return
	}

	var existingUser models.Users
	if err := database.Where("email =?", registerRequest.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	registerRequest.Password = string(hashedPassword)

	user := models.Users{
		Full_Name:     registerRequest.Full_Name,
		GenderInfo:    models.Gender(registerRequest.GenderInfo),
		ContactNumber: registerRequest.ContactNumber,
		Email:         registerRequest.Email,
		Password:      registerRequest.Password,
		Region:        registerRequest.Region,
		User_type:     models.Buyer,
	}
	if err := database.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}

func UserLogin(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Region   string `json:"region"`
	}
	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database := db.GetDB(loginRequest.Region)
	if database == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid region"})
		return
	}
	if database == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid region"})
		return
	}

	var admin models.Users
	if err := database.Where("email = ?", loginRequest.Email).First(&admin).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password"})
		return
	}

	jwtToken, err := utils.GenerateJWT(uint(admin.ID), admin.Email, string(admin.User_type), admin.Region)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Status": "Login successful", "token": jwtToken})
}
