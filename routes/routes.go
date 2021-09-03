package routes

import (
	"log"

	"github.com/ahmadirfaan/loan-payment/controller"
	"github.com/ahmadirfaan/loan-payment/repository"
	"github.com/ahmadirfaan/loan-payment/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, app *fiber.App) {
	customerRepository := repository.NewCustomerRepository(db)
	if err := customerRepository.Migrate(); err != nil {
		log.Fatal("Customer migrate err", err)
	}

	customerService := service.NewCustomerService(customerRepository)
	customerController := controller.NewCustomerController(customerService)
	app.Post("customer", customerController.AddCustomer)
}