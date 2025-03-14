package main

import (
	"fmt"

	"github.com/adityjoshi/E-Commerce-/service/authService/db"
)

func main() {
	fmt.Print("auth service")

	db.InitDB()
}
