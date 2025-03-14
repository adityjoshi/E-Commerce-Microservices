package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adityjoshi/E-Commerce-/service/authService/db"
	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Print("auth service")
	db.InitDB()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server := &http.Server{
		Addr:    ":8001",
		Handler: router,
	}

	log.Println("Server is running at :8001...")

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
