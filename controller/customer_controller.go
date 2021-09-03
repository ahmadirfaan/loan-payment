package controller

import (
	"fmt"
	"log"

	"github.com/ahmadirfaan/loan-payment/models"
	"github.com/ahmadirfaan/loan-payment/service"
	"github.com/gofiber/fiber/v2"
)

type CustomerController interface {
	AddCustomer(c *fiber.Ctx) error
}

type customerController struct {
	customerService service.CustomerService
}

//NewCustomerController -> returns new customer controller
func NewCustomerController(s service.CustomerService) CustomerController {
	return customerController{
		customerService: s,
	}
}

func (cs customerController) AddCustomer(c *fiber.Ctx) error {
	log.Print("[CustomerController]...add Customer")
	var customer models.Customer
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	customer, err := cs.customerService.Save(customer)
	fmt.Println(err)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": nil,
		"data":  customer,
	})
}
