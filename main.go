package main

import (
	"github.com/ahmadirfaan/loan-payment/config"
	"github.com/ahmadirfaan/loan-payment/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	//connect database
	db, _ := config.Connect()
	app := fiber.New()
	routes.SetupRoutes(db,app)
	app.Listen(":8000")
}