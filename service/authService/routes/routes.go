package routes

import (
	"github.com/adityjoshi/E-Commerce-Microservices/service/authService/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingReq *gin.Engine) {
	incomingReq.POST("/UserRegister", handlers.UserRegister)
	incomingReq.POST("/UserLogin", handlers.UserLogin)
}
